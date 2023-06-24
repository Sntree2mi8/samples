package service

import (
	"context"
	"database/sql"
	"github.com/DATA-DOG/go-txdb"
	"github.com/Sntree2mi8/samples/omiyage_management/omiyage/internal/core/usecase/command"
	"github.com/Sntree2mi8/samples/omiyage_management/omiyage/internal/infrastructure/mysql/gen/user"
	"github.com/Sntree2mi8/samples/omiyage_management/omiyage/internal/repository"
	omiyagev1 "github.com/Sntree2mi8/samples/omiyage_management/proto/gen/go/omiyage/v1"
	userv1 "github.com/Sntree2mi8/samples/omiyage_management/proto/gen/go/types/user/v1"
	"github.com/go-sql-driver/mysql"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/test/bufconn"
	"google.golang.org/protobuf/runtime/protoimpl"
	"log"
	"net"
	"testing"
)

// テストを書いていく順番.
// validationをなるべく早めに書いていくと, データ取得, validation, データ保存, responseの流れが綺麗になりやすい.

// assertを完結に描きたいので, testifyは使ってる.

// interceptorとかも含めてテストすることができるのでお便利.
func HelperNewOmiyageServiceBuffClient(t *testing.T, server omiyagev1.OmiyageServiceServer) omiyagev1.OmiyageServiceClient {
	t.Helper()

	lis := bufconn.Listen(1024 * 1024)
	s := grpc.NewServer()
	omiyagev1.RegisterOmiyageServiceServer(s, server)

	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatal(err)
		}
	}()

	t.Cleanup(func() {
		s.Stop()
	})

	conn, err := grpc.DialContext(
		context.Background(),
		"bufnet",
		grpc.WithContextDialer(func(ctx context.Context, s string) (net.Conn, error) {
			return lis.Dial()
		}),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		conn.Close()
	})

	return omiyagev1.NewOmiyageServiceClient(conn)
}

func HelperNewTestTXedDB(t *testing.T) *sql.DB {
	t.Helper()

	db, err := sql.Open("txdb", "identifier")
	if err != nil {
		t.Fatal("failed to setup test db")
	}

	return db
}

func TestOmiyageService_SignUp(t *testing.T) {
	config := mysql.Config{
		User:                 "user",
		Passwd:               "pass",
		Addr:                 "localhost:3306",
		DBName:               "omiyage",
		AllowNativePasswords: true,
		ParseTime:            true,
	}
	txdb.Register("txdb", "mysql", config.FormatDSN())

	type args struct {
		request *omiyagev1.SignUpRequest
	}
	tests := []struct {
		name           string
		args           args
		arrangeFunc    func(t *testing.T, db *sql.DB)
		want           *omiyagev1.SignUpResponse
		assertErrFunc  func(t *testing.T, err error)
		assertDataFunc func(t *testing.T, db *sql.DB, got *omiyagev1.SignUpResponse)
	}{
		// HappyPath
		{
			name: "SignUp成功",
			args: args{
				request: &omiyagev1.SignUpRequest{
					UserName: "UserName",
				},
			},
			want: &omiyagev1.SignUpResponse{
				Error: nil,
				User: &userv1.User{
					Name: "UserName",
				},
			},
			assertErrFunc: func(t *testing.T, err error) {
				t.Helper()

				assert.NoError(t, err)
			},
		},

		// 判断
		{
			name: "名前の入力は必須",
			args: args{
				request: &omiyagev1.SignUpRequest{
					UserName: "",
				},
			},
			want: &omiyagev1.SignUpResponse{
				Error: &omiyagev1.SignUpResponse_DomainError{
					Code:        omiyagev1.SignUpResponse_DomainError_DOMAIN_ERROR_CODE_VALIDATION_ERROR,
					Description: "",
				},
				User: nil,
			},
			assertErrFunc: func(t *testing.T, err error) {
				t.Helper()

				assert.NoError(t, err)
			},
		},

		// 保存の仕方
		{
			name: "保存できてる",
			args: args{
				request: &omiyagev1.SignUpRequest{
					UserName: "UserName",
				},
			},
			want: &omiyagev1.SignUpResponse{
				Error: nil,
				User: &userv1.User{
					Name: "UserName",
				},
			},
			assertErrFunc: func(t *testing.T, err error) {
				t.Helper()

				assert.NoError(t, err)
			},
			assertDataFunc: func(t *testing.T, db *sql.DB, got *omiyagev1.SignUpResponse) {
				t.Helper()

				res, err := user.New(db).FindUser(context.Background(), got.GetUser().GetId())
				if err != nil {
					t.Error(err)
				} else {
					assert.Equal(t, res.Name, "UserName")
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			omiyageDB := HelperNewTestTXedDB(t)
			defer omiyageDB.Close()

			client := HelperNewOmiyageServiceBuffClient(
				t,
				NewOmiyageService(
					command.NewUserSignUpUseCase(
						repository.NewUserRepository(omiyageDB),
					),
				),
			)

			// action
			got, err := client.SignUp(context.Background(), tt.args.request)

			// arrange
			if tt.arrangeFunc != nil {
				tt.arrangeFunc(t, omiyageDB)
			}

			// assert
			opts := []cmp.Option{
				// proto全般的にこのfieldの比較はテストの関心外として良い
				cmpopts.IgnoreTypes(protoimpl.MessageState{}),
				cmpopts.IgnoreTypes(protoimpl.SizeCache(0)),
				cmpopts.IgnoreTypes(protoimpl.UnknownFields{}),

				// ここはSeedを入れないかぎり毎回違うのでテストで確認する価値が薄い
				cmpopts.IgnoreFields(userv1.User{}, "Id"),
			}
			if d := cmp.Diff(tt.want, got, opts...); len(d) != 0 {
				t.Errorf("differs: (-got +want)\n%s", d)
			}

			if tt.assertErrFunc != nil {
				tt.assertErrFunc(t, err)
			}
			if tt.assertDataFunc != nil {
				tt.assertDataFunc(t, omiyageDB, got)
			}
		})
	}
}
