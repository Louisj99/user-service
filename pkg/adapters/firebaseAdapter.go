package adapters

import (
	"context"
	"firebase.google.com/go"
	"firebase.google.com/go/auth"
	"google.golang.org/api/option"
	"log"
	"user-service/pkg/entities"
)

type FirebaseAdapter struct {
	App  *firebase.App
	Auth *auth.Client
	Ctx  context.Context
}

func NewFirebaseAdapter(ctx context.Context, serviceAccountKeyFilePath string) (*FirebaseAdapter, error) {
	opt := option.WithCredentialsFile(serviceAccountKeyFilePath)
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		log.Fatalf("error initializing app: %v", err)
		return nil, err
	}

	auth, err := app.Auth(ctx)
	if err != nil {
		log.Fatalf("error getting Auth client: %v", err)
		return nil, err
	}

	return &FirebaseAdapter{
		App:  app,
		Auth: auth,
		Ctx:  ctx,
	}, nil
}
func (fa *FirebaseAdapter) GetUser(email string) (entities.User, error) {
	userRecord, err := fa.Auth.GetUserByEmail(fa.Ctx, email)
	if err != nil {
		return entities.User{}, err
	}

	user := entities.User{
		ID:    userRecord.UID,
		Email: userRecord.Email,
	}

	return user, nil
}

// Placeholder method to satisfy the PlaceholderInterface
func (fa *FirebaseAdapter) Placeholder(ctx context.Context, placeholder string) error {
	// Implement the method here
	return nil
}
