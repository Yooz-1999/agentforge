package logic

import (
	"context"
	"errors"
	"net/mail"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/Yooz-1999/agentforge/apps/core-rpc/internal/model"
	"github.com/Yooz-1999/agentforge/apps/core-rpc/internal/svc"
	"github.com/Yooz-1999/agentforge/apps/core-rpc/pb"
	"github.com/Yooz-1999/agentforge/pkg/auth"
	"github.com/Yooz-1999/agentforge/pkg/constants"

	"github.com/go-sql-driver/mysql"
	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type RegisterUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterUserLogic {
	return &RegisterUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RegisterUserLogic) RegisterUser(in *pb.RegisterUserRequest) (*pb.RegisterUserResponse, error) {
	email, password, nickname, err := normalizeRegistrationInput(in.Email, in.Password, in.Nickname)
	if err != nil {
		return nil, err
	}

	existingUser, err := l.svcCtx.UserRepo.FindByEmail(l.ctx, email)
	if err != nil {
		l.Errorf("find user by email failed: %v", err)
		return nil, status.Error(codes.Internal, "failed to query user")
	}
	if existingUser != nil {
		return nil, status.Error(codes.AlreadyExists, "email already exists")
	}

	passwordHash, err := auth.HashPassword(password)
	if err != nil {
		l.Errorf("hash password failed: %v", err)
		return nil, status.Error(codes.Internal, "failed to hash password")
	}

	now := time.Now()
	user := &model.User{
		Email:        email,
		PasswordHash: passwordHash,
		Nickname:     nickname,
		Status:       constants.UserStatusActive,
		CreatedAt:    now,
		UpdatedAt:    now,
	}

	if err := l.svcCtx.UserRepo.Create(l.ctx, user); err != nil {
		l.Errorf("create user failed: %v", err)
		var mysqlErr *mysql.MySQLError
		if errors.As(err, &mysqlErr) && mysqlErr.Number == 1062 {
			return nil, status.Error(codes.AlreadyExists, "email already exists")
		}

		return nil, status.Error(codes.Internal, "failed to create user")
	}

	return &pb.RegisterUserResponse{
		User: &pb.User{
			Id:        user.ID,
			Email:     user.Email,
			Nickname:  user.Nickname,
			Status:    user.Status,
			CreatedAt: user.CreatedAt.Format(time.RFC3339),
			UpdatedAt: user.UpdatedAt.Format(time.RFC3339),
		},
	}, nil
}

func normalizeRegistrationInput(emailInput, passwordInput, nicknameInput string) (string, string, string, error) {
	email := strings.TrimSpace(strings.ToLower(emailInput))
	nickname := strings.TrimSpace(nicknameInput)
	password := passwordInput

	if email == "" || nickname == "" || password == "" {
		return "", "", "", status.Error(codes.InvalidArgument, "email, password and nickname are required")
	}
	if utf8.RuneCountInString(email) > constants.EmailMaxCharacters {
		return "", "", "", status.Error(codes.InvalidArgument, "email is too long")
	}

	parsedAddress, err := mail.ParseAddress(email)
	if err != nil || parsedAddress.Name != "" || parsedAddress.Address != email {
		return "", "", "", status.Error(codes.InvalidArgument, "invalid email")
	}
	if utf8.RuneCountInString(password) < constants.PasswordMinCharacters {
		return "", "", "", status.Error(codes.InvalidArgument, "password must be at least 6 characters")
	}
	if len(password) > constants.PasswordMaxBytes {
		return "", "", "", status.Error(codes.InvalidArgument, "password is too long")
	}
	if utf8.RuneCountInString(nickname) > constants.NicknameMaxCharacters {
		return "", "", "", status.Error(codes.InvalidArgument, "nickname is too long")
	}

	return email, password, nickname, nil
}
