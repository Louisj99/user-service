package adapters

import (
	"context"
	"firebase.google.com/go"
	"firebase.google.com/go/auth"
	"google.golang.org/api/option"
	"log"
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

func (fa *FirebaseAdapter) CreateUser(email, password string) (*auth.UserRecord, error) {
	params := (&auth.UserToCreate{}).Email(email).Password(password)
	return fa.Auth.CreateUser(fa.Ctx, params)
}

func (fa *FirebaseAdapter) VerifyIDToken(idToken string) (*auth.Token, error) {
	return fa.Auth.VerifyIDToken(fa.Ctx, idToken)
}

// Placeholder method to satisfy the PlaceholderInterface
func (fa *FirebaseAdapter) Placeholder(ctx context.Context, placeholder string) error {
	// Implement the method here
	return nil
}
