package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"math/big"
	"os"
	_agency "server/agency"
	_agent "server/agent"
	"server/dataloaders"
	_email "server/email"
	_enterprise "server/enterprise"
	"server/finance"
	firebase "server/firebase"
	"server/graph/generated"
	"server/graph/model"
	"server/middleware"
	_notification "server/notification"
	_qrcode "server/qrcode"
	"server/reporttransaction"
	"server/transaction"
	_user "server/user"
	"server/utils"
	"sync"

	"github.com/99designs/gqlgen/graphql"
)

func init() {
	utils.LoadEnv()
}

type TransactionObserver struct {
	ID      string
	Payment chan *model.Paiement
}
type NotificationOberser struct {
	UserID       string
	Notification chan *model.Notification
}
type Payment struct {
	ID        string
	Observers sync.Map
}
type Notif struct {
	ID        string
	Observers sync.Map
}

func (r *Resolver) GetPayment(id string) *Payment {
	payment, _ := r.Payments.LoadOrStore(id, &Payment{
		ID:        id,
		Observers: sync.Map{},
	})
	return payment.(*Payment)
}
func (r *Resolver) GetNotification(id string) *Payment {
	payment, _ := r.Payments.LoadOrStore(id, &Notif{
		ID:        id,
		Observers: sync.Map{},
	})
	return payment.(*Payment)
}

func (r *mutationResolver) Connect(ctx context.Context, token string) (string, error) {
	authToken, err := firebase.Connect().VerifyIdToken(ctx, token)
	if err != nil {
		return "", err
	}

	customToken, err := firebase.Connect().CreateCustomToken(ctx, authToken.UID)
	if err != nil {
		return "", err
	}

	return customToken, nil
}

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, user *model.UserInput) (*model.UserCreated, error) {
	key := utils.UIDCtxKey
	x := utils.ForContextUser(ctx, key)
	ip := middleware.ForRemoteAddressContext(ctx)
	return _user.CreateUser(ctx, user, x, ip)
}

// DeleteUser is the resolver for the deleteUser field.
func (r *mutationResolver) DeleteUser(ctx context.Context) (bool, error) {
	key := utils.UserCtxKey
	x := utils.ForContextDBUser(ctx, key)
	ip := middleware.ForRemoteAddressContext(ctx)
	return _user.DeleteUser(*x, ip)
}

// SendEmail is the resolver for the sendEmail field.
func (r *mutationResolver) SendEmail(ctx context.Context, email *model.EmailInput) (bool, error) {
	ip := middleware.ForRemoteAddressContext(ctx)
	return _email.SendEmail(email, ip)
}

// SuscribeToNewsLetter is the resolver for the suscribeToNewsLetter field.
func (r *mutationResolver) SuscribeToNewsLetter(ctx context.Context, email string) (bool, error) {
	ip := middleware.ForRemoteAddressContext(ctx)
	return _email.SuscribeToNewsLetter(email, ip)
}

// CreateTransfer is the resolver for the createTransfer field.
func (r *mutationResolver) CreateTransfer(ctx context.Context, address *string, token string, amount float64, pinCode string, destinationUser string) (bool, error) {
	key := utils.UserCtxKey
	x := utils.ForContextDBUser(ctx, key)
	__pincodeUser, err := utils.VerifyPincode(*x, pinCode)
	if err != nil {
		return false, err
	}

	ip := middleware.ForRemoteAddressContext(ctx)
	success, err := transaction.CreateTransfer(ctx, address, os.Getenv("DEFAULT_CURRENCY"), amount, destinationUser, *__pincodeUser, ip)
	return *success, err
}

// UpdateProfilePicture is the resolver for the updateProfilePicture field.
func (r *mutationResolver) UpdateProfilePicture(ctx context.Context, link string) (bool, error) {
	key := utils.UserCtxKey
	x := utils.ForContextDBUser(ctx, key)
	ip := middleware.ForRemoteAddressContext(ctx)
	return _user.UpdateProfilPicture(*x, link, ip)
}

// CreateContact is the resolver for the createContact field.
func (r *mutationResolver) CreateContact(ctx context.Context, contact string) (bool, error) {
	key := utils.UserCtxKey
	x := utils.ForContextDBUser(ctx, key)
	ip := middleware.ForRemoteAddressContext(ctx)
	success, err := _user.AddContact(*x, contact, ip)
	return *success, err
}

// RemoveContact is the resolver for the removeContact field.
func (r *mutationResolver) RemoveContact(ctx context.Context, contactID string) (bool, error) {
	key := utils.UserCtxKey
	x := utils.ForContextDBUser(ctx, key)
	ip := middleware.ForRemoteAddressContext(ctx)
	success, err := _user.RemoveContact(*x, contactID, ip)
	return *success, err
}

// AddAGency is the resolver for the addAGency field.
func (r *mutationResolver) AddAGency(ctx context.Context, agency model.AgencyInpyt, pinCode string) (string, error) {
	key := utils.UserCtxKey
	x := utils.ForContextDBUser(ctx, key)
	__pincodeUser, err := utils.VerifyPincode(*x, pinCode)
	if err != nil {
		return "", err
	}
	ip := middleware.ForRemoteAddressContext(ctx)
	return _agency.CreateAgency(ctx, agency, __pincodeUser.ID, ip)
}

// AddWithDraw is the resolver for the addWithDraw field.
func (r *mutationResolver) AddWithDraw(ctx context.Context, withdraw model.WithdrawInput, pinCode string) (string, error) {
	key := utils.UserCtxKey
	x := utils.ForContextDBUser(ctx, key)
	__pincodeUser, err := utils.VerifyPincode(*x, pinCode)
	if err != nil {
		return "", err
	}
	ip := middleware.ForRemoteAddressContext(ctx)
	return transaction.AddWithdraw(ctx, withdraw, pinCode, *__pincodeUser, ip)
}

// AddTopUp is the resolver for the addTopUp field.
func (r *mutationResolver) AddTopUp(ctx context.Context, topup model.TopUpInput, pinCode string) (string, error) {
	key := utils.UserCtxKey
	x := utils.ForContextDBUser(ctx, key)
	__pincodeUser, err := utils.VerifyPincode(*x, pinCode)
	if err != nil {
		return "", err
	}
	ip := middleware.ForRemoteAddressContext(ctx)
	return transaction.AddTopup(ctx, topup, pinCode, *__pincodeUser, ip)
}

// ConfirmTransactionAgent is the resolver for the confirmTransactionAgent field.
func (r *mutationResolver) ConfirmTransactionAgent(ctx context.Context, transactionID string, typeArg model.PaymentType, token string, pinCode string) (bool, error) {
	key := utils.UserCtxKey
	x := utils.ForContextDBUser(ctx, key)
	__pincodeUser, err := utils.VerifyPincode(*x, pinCode)
	if err != nil {
		return false, err
	}
	ip := middleware.ForRemoteAddressContext(ctx)
	return _agent.ConfirmTransaction(ctx, transactionID, typeArg, pinCode, *__pincodeUser, ip)
}

// CancelTransactionAgent is the resolver for the cancelTransactionAgent field.
func (r *mutationResolver) CancelTransactionAgent(ctx context.Context, transactionID *string, typeArg *model.PaymentType, pinCode string) (bool, error) {
	key := utils.UserCtxKey
	x := utils.ForContextDBUser(ctx, key)
	__pincodeUser, err := utils.VerifyPincode(*x, pinCode)
	if err != nil {
		return false, err
	}
	ip := middleware.ForRemoteAddressContext(ctx)
	return _agent.CancelTransaction(ctx, *transactionID, *typeArg, pinCode, *__pincodeUser, ip)
}

// CancelTransactionUser is the resolver for the cancelTransactionUser field.
func (r *mutationResolver) CancelTransactionUser(ctx context.Context, transactionID *string, typeArg *model.PaymentType, pinCode string) (bool, error) {
	key := utils.UserCtxKey
	x := utils.ForContextDBUser(ctx, key)
	__pincodeUser, err := utils.VerifyPincode(*x, pinCode)
	if err != nil {
		return false, err
	}
	ip := middleware.ForRemoteAddressContext(ctx)
	return transaction.CancelTransactionUser(ctx, *transactionID, *typeArg, pinCode, *__pincodeUser, ip)
}

// SetIsOnline is the resolver for the setIsOnline field.
func (r *mutationResolver) SetIsOnline(ctx context.Context, toggle *bool) (bool, error) {
	key := utils.UserCtxKey
	x := utils.ForContextDBUser(ctx, key)
	ip := middleware.ForRemoteAddressContext(ctx)
	return _user.ToggleOnlineStatus(*x, *toggle, ip)
}

// UpdateFcmToken is the resolver for the updateFcmToken field.
func (r *mutationResolver) UpdateFcmToken(ctx context.Context, fcmToken *string) (bool, error) {
	key := utils.UserCtxKey
	x := utils.ForContextDBUser(ctx, key)
	ip := middleware.ForRemoteAddressContext(ctx)
	return _user.UpdateFcmToken(x, fcmToken, ip)
}

// ChangePinCode is the resolver for the changePinCode field.
func (r *mutationResolver) ChangePinCode(ctx context.Context, newPin *string) (bool, error) {
	key := utils.UserCtxKey
	x := utils.ForContextDBUser(ctx, key)
	ip := middleware.ForRemoteAddressContext(ctx)
	return _user.ChangePinCode(*x, *newPin, ip)
}

// SetAllNotificationToRead is the resolver for the setAllNotificationToRead field.
func (r *mutationResolver) SetAllNotificationToRead(ctx context.Context) (bool, error) {
	key := utils.UserCtxKey
	x := utils.ForContextDBUser(ctx, key)
	return _notification.SetAllNotificationToRead(ctx, x.ID)
}

// UploadFile is the resolver for the uploadFile field.
func (r *mutationResolver) UploadFile(ctx context.Context, file graphql.Upload, typeArg string) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

// ReportTransaction is the resolver for the reportTransaction field.
func (r *mutationResolver) ReportTransaction(ctx context.Context, transactionID string, message string) (*bool, error) {
	key := utils.UserCtxKey
	x := utils.ForContextDBUser(ctx, key)
	ip := middleware.ForRemoteAddressContext(ctx)
	return reporttransaction.ReportTransaction(*x, transactionID, message, ip)
}

// AssignRole is the resolver for the assignRole field.
func (r *mutationResolver) AssignRole(ctx context.Context, userID string, role string, pinCode string, token string) (*bool, error) {
	key := utils.UserCtxKey
	x := utils.ForContextDBUser(ctx, key)
	__pincodeUser, err := utils.VerifyPincode(*x, pinCode)
	if err != nil {
		return nil, err
	}
	ip := middleware.ForRemoteAddressContext(ctx)
	return _agent.AssignRoleToAddress(ctx, userID, role, *__pincodeUser, ip)
}

// UnassignRole is the resolver for the unassignRole field.
func (r *mutationResolver) UnassignRole(ctx context.Context, userID string, role string, pinCode string, token string) (*bool, error) {
	key := utils.UserCtxKey
	x := utils.ForContextDBUser(ctx, key)
	__pincodeUser, err := utils.VerifyPincode(*x, pinCode)
	if err != nil {
		return nil, err
	}
	ip := middleware.ForRemoteAddressContext(ctx)
	return _agent.UnAssignRoleToAddress(ctx, userID, role, *__pincodeUser, ip)
}

// CreateEnterprise is the resolver for the createEnterprise field.
func (r *mutationResolver) CreateEnterprise(ctx context.Context, enterprise model.EnterpriseInput) (*model.Enterprise, error) {
	key := utils.UserCtxKey
	x := utils.ForContextDBUser(ctx, key)
	ip := middleware.ForRemoteAddressContext(ctx)
	return _enterprise.CreateEnterprise(ctx, enterprise, *x, ip)
}

// AuthenticateForPayment is the resolver for the authenticateForPayment field.
func (r *mutationResolver) AuthenticateForPayment(ctx context.Context, amount float64, ref *string) (model.QRCodeOwner, error) {
	key := utils.EnterpriseCtxKey
	x := utils.ForContextDBEnnterprise(ctx, key)
	ip := middleware.ForRemoteAddressContext(ctx)
	return transaction.AuthenticateForTransaction(ctx, amount, ref, x, ip)
}

// RecreateEnterprisePublishableKey is the resolver for the recreateEnterprisePublishableKey field.
func (r *mutationResolver) RecreateEnterprisePublishableKey(ctx context.Context, enterpriseID string, pinCode string) ([]*model.Enterprise, error) {
	key := utils.UserCtxKey
	x := utils.ForContextDBUser(ctx, key)
	__pincodeUser, err := utils.VerifyPincode(*x, pinCode)
	if err != nil {
		return nil, err
	}

	ip := middleware.ForRemoteAddressContext(ctx)
	return _enterprise.RecreatePublishableKey(ctx, enterpriseID, __pincodeUser.ID, ip)
}

// RecreateEnterprisePrivateKey is the resolver for the recreateEnterprisePrivateKey field.
func (r *mutationResolver) RecreateEnterprisePrivateKey(ctx context.Context, enterpriseID string, pinCode string) ([]*model.Enterprise, error) {
	key := utils.UserCtxKey
	x := utils.ForContextDBUser(ctx, key)
	__pincodeUser, err := utils.VerifyPincode(*x, pinCode)
	if err != nil {
		return nil, err
	}
	ip := middleware.ForRemoteAddressContext(ctx)
	return _enterprise.RcreatePrivateKey(ctx, enterpriseID, __pincodeUser.ID, ip)
}

// RemoveEnterprise is the resolver for the removeEnterprise field.
func (r *mutationResolver) RemoveEnterprise(ctx context.Context, enterpriseID string, pinCode string) ([]*model.Enterprise, error) {
	key := utils.UserCtxKey
	x := utils.ForContextDBUser(ctx, key)
	__pincodeUser, err := utils.VerifyPincode(*x, pinCode)
	if err != nil {
		return nil, err
	}
	ip := middleware.ForRemoteAddressContext(ctx)
	return _enterprise.RemoveEnterprise(ctx, enterpriseID, __pincodeUser.ID, ip)
}

// ChangeDefaultEnterprise is the resolver for the changeDefaultEnterprise field.
func (r *mutationResolver) ChangeDefaultEnterprise(ctx context.Context, enterpriseID string) ([]*model.Enterprise, error) {
	key := utils.UserCtxKey
	x := utils.ForContextDBUser(ctx, key)
	ip := middleware.ForRemoteAddressContext(ctx)
	return _enterprise.ChangeDefaultEnterprise(ctx, enterpriseID, x.ID, ip)
}

// UpdateEnterpriseType is the resolver for the updateEnterpriseType field.
func (r *mutationResolver) UpdateEnterpriseType(ctx context.Context, enterpriseID string, typeArg string, country string) ([]*model.Enterprise, error) {
	key := utils.UserCtxKey
	x := utils.ForContextDBUser(ctx, key)
	ip := middleware.ForRemoteAddressContext(ctx)
	return _enterprise.UpdateEnterpriseType(ctx, enterpriseID, x.ID, typeArg, country, ip)
}

// UpdatePersonnalInformation is the resolver for the updatePersonnalInformation field.
func (r *mutationResolver) UpdatePersonnalInformation(ctx context.Context, enterpriseID string, firstName string, lastName string, email string, address string, city string, state string, zip string) ([]*model.Enterprise, error) {
	key := utils.UserCtxKey
	x := utils.ForContextDBUser(ctx, key)
	ip := middleware.ForRemoteAddressContext(ctx)
	return _enterprise.UpdatePersonnalInformation(ctx, enterpriseID, x.ID, firstName, lastName, email, address, city, state, zip, ip)
}

// UpdateEnterpriseInformation is the resolver for the updateEnterpriseInformation field.
func (r *mutationResolver) UpdateEnterpriseInformation(ctx context.Context, enterpriseID string, rccm string, sector string, website *string, description *string) ([]*model.Enterprise, error) {
	key := utils.UserCtxKey
	x := utils.ForContextDBUser(ctx, key)
	ip := middleware.ForRemoteAddressContext(ctx)
	return _enterprise.UpdateEnterpriseInformation(ctx, enterpriseID, x.ID, rccm, sector, website, description, ip)
}

// UpdateExecutionInformation is the resolver for the updateExecutionInformation field.
func (r *mutationResolver) UpdateExecutionInformation(ctx context.Context, enterpriseID string, sellingPyshicalGoods *bool, selfShipping *bool, shippingDelay *string) ([]*model.Enterprise, error) {
	key := utils.UserCtxKey
	x := utils.ForContextDBUser(ctx, key)
	ip := middleware.ForRemoteAddressContext(ctx)
	return _enterprise.UpdateExecutionInformation(ctx, enterpriseID, x.ID, sellingPyshicalGoods, selfShipping, shippingDelay, ip)
}

// UpdatePublicInformation is the resolver for the updatePublicInformation field.
func (r *mutationResolver) UpdatePublicInformation(ctx context.Context, enterpriseID string, name string, libelle string, libelleAbreged string, email *string, phone model.PhoneInput) ([]*model.Enterprise, error) {
	key := utils.UserCtxKey
	x := utils.ForContextDBUser(ctx, key)
	ip := middleware.ForRemoteAddressContext(ctx)
	return _enterprise.UpdatePublicInformation(ctx, enterpriseID, x.ID, name, libelle, libelleAbreged, email, phone, ip)
}

// PayUnConfirmedTransaction is the resolver for the payUnConfirmedTransaction field.
func (r *mutationResolver) PayUnConfirmedTransaction(ctx context.Context, enterpriseID string, pinCode string, transactionID string) (*model.Paiement, error) {
	key := utils.UserCtxKey
	x := utils.ForContextDBUser(ctx, key)
	__pincodedUser, err := utils.VerifyPincode(*x, pinCode)
	if err != nil {
		return nil, err
	}
	ip := middleware.ForRemoteAddressContext(ctx)
	res, err := _enterprise.PayUnConfirmedTransaction(ctx, enterpriseID, transactionID, *__pincodedUser, ip)
	if res != nil && err == nil {
		payment := r.GetPayment(transactionID)
		payment.Observers.Range(func(_, v interface{}) bool {
			observer := v.(*TransactionObserver)
			if observer.ID == x.ID {
				observer.Payment <- res
			}
			return true
		})
	}
	return res, err
}

// PayEnterprise is the resolver for the payEnterprise field.
func (r *mutationResolver) PayEnterprise(ctx context.Context, enterpriseID string, amount float64, pinCode string) (*model.Paiement, error) {
	key := utils.UserCtxKey
	x := utils.ForContextDBUser(ctx, key)
	__pincodeUser, err := utils.VerifyPincode(*x, pinCode)
	if err != nil {
		return nil, err
	}
	ip := middleware.ForRemoteAddressContext(ctx)
	return _enterprise.PayEnterprise(ctx, enterpriseID, amount, pinCode, *__pincodeUser, ip)
}

// RefundTransaction is the resolver for the refundTransaction field.
func (r *mutationResolver) RefundTransaction(ctx context.Context, enterpriseID string, pinCode string, transactionID string) (bool, error) {
	keyuser := utils.UserCtxKey
	db_user := utils.ForContextDBUser(ctx, keyuser)
	company, err := _enterprise.GetEnterpriseByIdWithUseriD(ctx, enterpriseID, db_user.ID)
	if err != nil {
		return false, err
	}

	__pincodeEnterprise, err := utils.VerifyPincodeWithEnterprise(*company, *db_user, pinCode)
	if err != nil {
		return false, err
	}

	ip := middleware.ForRemoteAddressContext(ctx)
	success, err := _enterprise.RefundTransaction(ctx, enterpriseID, transactionID, *__pincodeEnterprise, ip)
	return *success, err
}

// CancelTransactionEnterprise is the resolver for the cancelTransactionEnterprise field.
func (r *mutationResolver) CancelTransactionEnterprise(ctx context.Context, enterpriseID string, pinCode string, transactionID string) (bool, error) {
	key := utils.UserCtxKey
	x := utils.ForContextDBUser(ctx, key)
	_, err := utils.VerifyPincode(*x, pinCode)
	if err != nil {
		return false, err
	}

	ip := middleware.ForRemoteAddressContext(ctx)
	success, err := _enterprise.CancelEnterpriseTransaction(ctx, enterpriseID, transactionID, ip)
	return *success, err
}

// TransferMoneyEnterprise is the resolver for the transferMoneyEnterprise field.
func (r *mutationResolver) TransferMoneyEnterprise(ctx context.Context, enterpriseID string, pinCode string, publicKey string, amount float64) (bool, error) {
	keyuser := utils.UserCtxKey
	db_user := utils.ForContextDBUser(ctx, keyuser)
	company, err := _enterprise.GetEnterpriseByIdWithUseriD(ctx, enterpriseID, db_user.ID)
	if err != nil {
		return false, err
	}
	__pincodeEnterprise, err := utils.VerifyPincodeWithEnterprise(*company, *db_user, pinCode)
	if err != nil {
		return false, err
	}

	ip := middleware.ForRemoteAddressContext(ctx)
	success, err := _enterprise.SendMoney(ctx, enterpriseID, pinCode, publicKey, amount, db_user.ID, __pincodeEnterprise, ip)
	return *success, err
}

// UserExist is the resolver for the userExist field.
func (r *queryResolver) UserExist(ctx context.Context, uid string) (bool, error) {
	return _user.UserExist(uid)
}

// UsersExist is the resolver for the usersExist field.
func (r *queryResolver) UsersExist(ctx context.Context) (*model.User, error) {
	key := utils.UserCtxKey
	x := utils.ForContextDBUser(ctx, key)
	if x != nil {
		return x, nil
	}
	return nil, fmt.Errorf("user dosnt exist in our database")
}

// LoadNotification is the resolver for the loadNotification field.
func (r *queryResolver) LoadNotification(ctx context.Context) ([]*model.Notification, error) {
	key := utils.UserCtxKey
	x := utils.ForContextDBUser(ctx, key)
	return _notification.LoadNotification(ctx, x.ID)
}

// LoadNotificationCount is the resolver for the loadNotificationCount field.
func (r *queryResolver) LoadNotificationCount(ctx context.Context) (float64, error) {
	key := utils.UserCtxKey
	x := utils.ForContextDBUser(ctx, key)
	count, err := _notification.LoadNotificationCount(ctx, x.ID)
	return *count, err
}

// LoadBalance is the resolver for the loadBalance field.
func (r *queryResolver) LoadBalance(ctx context.Context) (*model.Wallet, error) {
	key := utils.UserCtxKey
	x := utils.ForContextDBUser(ctx, key)

	return _user.LoadBalance(x.Keypair.PublicKey)
}

// SearchUser is the resolver for the searchUser field.
func (r *queryResolver) SearchUser(ctx context.Context, searchText *string) ([]*model.UserSmall, error) {
	key := utils.UserCtxKey
	x := utils.ForContextDBUser(ctx, key)
	contacts, err := _user.SearchUserUnAdded(x, *searchText)
	return contacts, err
}

// GetAllUserContact is the resolver for the getAllUserContact field.
func (r *queryResolver) GetAllUserContact(ctx context.Context, searchText string) ([]*model.UserSmall, error) {
	key := utils.UserCtxKey
	x := utils.ForContextDBUser(ctx, key)
	return _user.SearchUser(x, searchText)
}

// GetAllContactNotAdded is the resolver for the getAllContactNotAdded field.
func (r *queryResolver) GetAllContactNotAdded(ctx context.Context, searchText string) ([]*model.UserSmall, error) {
	key := utils.UserCtxKey
	x := utils.ForContextDBUser(ctx, key)
	contacts, err := _user.SearchUserUnAdded(x, searchText)
	return contacts, err
}

// GetActivity is the resolver for the getActivity field.
func (r *queryResolver) GetActivity(ctx context.Context) ([]*model.Paiement, error) {
	key := utils.UserCtxKey
	x := utils.ForContextDBUser(ctx, key)
	return _user.LoadAllActivities(x)
}

// LoadQRCode is the resolver for the loadQRCode field.
func (r *queryResolver) LoadQRCode(ctx context.Context) (string, error) {
	key := utils.UserCtxKey
	x := utils.ForContextDBUser(ctx, key)
	return x.Keypair.PublicKey, nil
}

// LoadTokenSupply is the resolver for the loadTokenSupply field.
func (r *queryResolver) LoadTokenSupply(ctx context.Context, token string) (int, error) {
	supply, err := finance.LoadTokenSupply()
	return int(supply.Int64()), err
}

// RetrieveAllAgnecies is the resolver for the retrieveAllAgnecies field.
func (r *queryResolver) RetrieveAllAgnecies(ctx context.Context) ([]*model.Agency, error) {
	return _agency.RetrieveAllAgencies(ctx)
}

// RetrieveAllAgenciesTransactions is the resolver for the retrieveAllAgenciesTransactions field.
func (r *queryResolver) RetrieveAllAgenciesTransactions(ctx context.Context) ([]*model.Paiement, error) {
	panic(fmt.Errorf("not implemented"))
}

// GetAllParticipatingTransactions is the resolver for the getAllParticipatingTransactions field.
func (r *queryResolver) GetAllParticipatingTransactions(ctx context.Context) ([]*model.Paiement, error) {
	key := utils.UserCtxKey
	x := utils.ForContextDBUser(ctx, key)
	return _user.GetAllParticipatingTransactions(*x)
}

// GetQROwner is the resolver for the getQrOwner field.
func (r *queryResolver) GetQROwner(ctx context.Context, qrcode string) (model.QRCodeOwner, error) {
	return _qrcode.QueryQrCodeUsers(qrcode)
}

// GetTransactionByIDUnauthed is the resolver for the getTransactionByIdUnauthed field.
func (r *queryResolver) GetTransactionByIDUnauthed(ctx context.Context, id string) (*model.Paiement, error) {
	return transaction.GetTransactionByIdUnauthed(ctx, id)
}

// GetTransactionByID is the resolver for the getTransactionById field.
func (r *queryResolver) GetTransactionByID(ctx context.Context, id string) (*model.Paiement, error) {
	key := utils.UserCtxKey
	x := utils.ForContextDBUser(ctx, key)
	return transaction.GetTransactionById(ctx, id, x.ID)
}

// GetTransactionByIDAgent is the resolver for the getTransactionByIdAgent field.
func (r *queryResolver) GetTransactionByIDAgent(ctx context.Context, id string) (*model.Paiement, error) {
	return transaction.GetTransactionByIdAgent(ctx, id)
}

// GetselfEmployedPDF is the resolver for the getselfEmployedPDF field.
func (r *queryResolver) GetselfEmployedPDF(ctx context.Context) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

// GetAllUserEnterprise is the resolver for the getAllUserEnterprise field.
func (r *queryResolver) GetAllUserEnterprise(ctx context.Context) ([]*model.Enterprise, error) {
	key := utils.UserCtxKey
	x := utils.ForContextDBUser(ctx, key)
	return _enterprise.GetAllEnterpriseForAUser(ctx, x.ID)
}

// GetAllTransactionByEnterpriseID is the resolver for the getAllTransactionByEnterpriseId field.
func (r *queryResolver) GetAllTransactionByEnterpriseID(ctx context.Context, enterpriseID string, from string, to string, limit float64, skip float64) (*model.TransactionWithPageInfo, error) {
	key := utils.UserCtxKey
	x := utils.ForContextDBUser(ctx, key)
	return transaction.GetAllTransactionByEnterpriseId(ctx, enterpriseID, from, to, limit, skip, x.ID)
}

// GetSuccessFullTransactionByEnterpriseID is the resolver for the getSuccessFullTransactionByEnterpriseId field.
func (r *queryResolver) GetSuccessFullTransactionByEnterpriseID(ctx context.Context, enterpriseID string, from string, to string, limit float64, skip float64) (*model.TransactionWithPageInfo, error) {
	key := utils.UserCtxKey
	x := utils.ForContextDBUser(ctx, key)
	return transaction.GetSuccessFullTransactionByEnterpriseId(enterpriseID, from, to, limit, skip, x.ID)
}

// GetRefundedTransactionByEnterpriseID is the resolver for the getRefundedTransactionByEnterpriseId field.
func (r *queryResolver) GetRefundedTransactionByEnterpriseID(ctx context.Context, enterpriseID string, from string, to string, limit float64, skip float64) (*model.TransactionWithPageInfo, error) {
	key := utils.UserCtxKey
	x := utils.ForContextDBUser(ctx, key)
	return transaction.GetRefundedTransactionByEnterpriseId(enterpriseID, from, to, limit, skip, x.ID)
}

// GetNonCapturedTransactionByEnterpriseID is the resolver for the getNonCapturedTransactionByEnterpriseId field.
func (r *queryResolver) GetNonCapturedTransactionByEnterpriseID(ctx context.Context, enterpriseID string, from string, to string, limit float64, skip float64) (*model.TransactionWithPageInfo, error) {
	key := utils.UserCtxKey
	x := utils.ForContextDBUser(ctx, key)
	return transaction.GetNonCapturedTransactionByEnterpriseId(enterpriseID, from, to, limit, skip, x.ID)
}

// GetFailedTransactionByEnterpriseID is the resolver for the getFailedTransactionByEnterpriseId field.
func (r *queryResolver) GetFailedTransactionByEnterpriseID(ctx context.Context, enterpriseID string, from string, to string, limit float64, skip float64) (*model.TransactionWithPageInfo, error) {
	key := utils.UserCtxKey
	x := utils.ForContextDBUser(ctx, key)
	return transaction.GetFailedTransactionByEnterpriseId(enterpriseID, from, to, limit, skip, x.ID)
}

// GetProfilNetChartData is the resolver for the getProfilNetChartData field.
func (r *queryResolver) GetProfilNetChartData(ctx context.Context, enterpriseID string, from string, to string) (*model.ChartData, error) {
	key := utils.UserCtxKey
	x := utils.ForContextDBUser(ctx, key)
	return transaction.GetProfilNetChartData(enterpriseID, from, to, x.ID)
}

// GetProfilBrutChartData is the resolver for the getProfilBrutChartData field.
func (r *queryResolver) GetProfilBrutChartData(ctx context.Context, enterpriseID string, from string, to string) (*model.ChartData, error) {
	key := utils.UserCtxKey
	x := utils.ForContextDBUser(ctx, key)
	return transaction.GetProfilBrutChartData(enterpriseID, from, to, x.ID)
}

// GetProfilNonCarpturedChartData is the resolver for the getProfilNonCarpturedChartData field.
func (r *queryResolver) GetProfilNonCarpturedChartData(ctx context.Context, enterpriseID string, from string, to string) (*model.ChartData, error) {
	key := utils.UserCtxKey
	x := utils.ForContextDBUser(ctx, key)
	return transaction.GetProfilNonCarpturedChartData(enterpriseID, from, to, x.ID)
}

// GetEnterpriseBalance is the resolver for the getEnterpriseBalance field.
func (r *queryResolver) GetEnterpriseBalance(ctx context.Context, enterpriseID string) (float64, error) {
	key := utils.UserCtxKey
	x := utils.ForContextDBUser(ctx, key)
	balance := _enterprise.GetEnterpriseBalance(ctx, enterpriseID, x.ID)
	return *balance, nil
}

// GetEntrepriseByID is the resolver for the getEntrepriseById field.
func (r *queryResolver) GetEntrepriseByID(ctx context.Context, enterpriseID string) (*model.Enterprise, error) {
	return _enterprise.GetEnterpriseById(ctx, enterpriseID)
}

// GetPDF is the resolver for the getPdf field.
func (r *queryResolver) GetPDF(ctx context.Context, enterpriseID string) (string, error) {
	key := utils.UserCtxKey
	x := utils.ForContextDBUser(ctx, key)
	pdf, err := _enterprise.GetEnterprisePDF(ctx, enterpriseID, *x)
	return pdf, err
}

func (r *notificationResolver) From(ctx context.Context, obj *model.Notification) (*model.UserSmall, error) {
	if obj.FromID == nil {
		return &model.UserSmall{}, nil
	}
	return dataloaders.NewLoaders().User(ctx, *obj.FromID)
}

func (r *paiementResolver) Agency(ctx context.Context, obj *model.Paiement) (*model.Agency, error) {
	if obj.AgencyID == nil {
		return &model.Agency{}, nil
	}
	return dataloaders.NewLoaders().Agency(ctx, *obj.AgencyID)
}

func (r *paiementResolver) Enterprise(ctx context.Context, obj *model.Paiement) (*model.EnterpriseSmall, error) {
	if obj.EnterpriseID == nil {
		return &model.EnterpriseSmall{}, nil
	}
	return dataloaders.NewLoaders().Enterprise(ctx, *obj.EnterpriseID)
}

func (r *paiementResolver) Creator(ctx context.Context, obj *model.Paiement) (*model.UserSmall, error) {
	if obj.CreatorID == nil {
		return nil, nil
	}

	user, err := dataloaders.NewLoaders().User(ctx, *obj.CreatorID)
	return user, err
}

func (r *paiementResolver) Cancellor(ctx context.Context, obj *model.Paiement) (*model.UserSmall, error) {
	if obj.CancellorID == nil {
		return &model.UserSmall{}, nil
	}
	return dataloaders.NewLoaders().User(ctx, *obj.CreatorID)
}

func (r *paiementResolver) Amount(ctx context.Context, obj *model.Paiement) (float64, error) {
	_amount, _ := new(big.Int).SetString(obj.AmountInt64, 0)
	amountbigFloat := finance.FromWei(*_amount)
	output, _ := amountbigFloat.Float64()
	return output, nil
}

func (r *paiementResolver) Fee(ctx context.Context, obj *model.Paiement) (*float64, error) {
	_amount, _ := new(big.Int).SetString(obj.FeeInt64, 0)
	amountbigFloat := finance.FromWei(*_amount)
	output, _ := amountbigFloat.Float64()
	return &output, nil
}

func (r *paiementResolver) FeeEnterprise(ctx context.Context, obj *model.Paiement) (*float64, error) {
	_amount, _ := new(big.Int).SetString(obj.FeeEnterpriseInt64, 0)
	amountbigFloat := finance.FromWei(*_amount)
	output, _ := amountbigFloat.Float64()
	return &output, nil
}

func (r *paiementResolver) Validator(ctx context.Context, obj *model.Paiement) (*model.UserSmall, error) {
	if obj.ValidatorID == nil {
		return nil, nil
	}
	return dataloaders.NewLoaders().User(ctx, *obj.ValidatorID)
}

func (r *paiementResolver) DestinationUser(ctx context.Context, obj *model.Paiement) (*model.UserSmall, error) {
	if obj.DestinationUserID == nil {
		return nil, nil
	}
	return dataloaders.NewLoaders().User(ctx, *obj.DestinationUserID)
}

func (r *enterpriseResolver) PublishableKey(ctx context.Context, obj *model.Enterprise) (string, error) {
	return utils.Ase256Decode(os.Getenv("SERVER_SECRET_KEY"), obj.PublishableKeyString), nil
}

func (r *enterpriseResolver) PrivateKey(ctx context.Context, obj *model.Enterprise) (string, error) {
	return utils.Ase256Decode(os.Getenv("SERVER_SECRET_KEY"), obj.PrivateKeystring), nil
}

// NotificationAdded is the resolver for the notificationAdded field.
func (r *subscriptionResolver) NotificationAdded(ctx context.Context, listener string) (<-chan *model.Notification, error) {
	panic(fmt.Errorf("not implemented"))
}

// TransactionPayed is the resolver for the transactionPayed field.
func (r *subscriptionResolver) TransactionPayed(ctx context.Context, id string) (<-chan *model.Paiement, error) {
	payment := r.GetPayment(id)
	events := make(chan *model.Paiement, 1)

	go func() {
		<-ctx.Done()
		payment.Observers.Delete(id)
	}()

	payment.Observers.Store(id, TransactionObserver{
		ID:      id,
		Payment: events,
	})

	return events, nil
}

// ContactAdded is the resolver for the contactAdded field.
func (r *subscriptionResolver) ContactAdded(ctx context.Context, listener string) (<-chan *model.Notification, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Subscription returns generated.SubscriptionResolver implementation.
func (r *Resolver) Subscription() generated.SubscriptionResolver { return &subscriptionResolver{r} }

func (r *Resolver) Paiement() generated.PaiementResolver { return &paiementResolver{r} }

func (r *Resolver) Notification() generated.NotificationResolver { return &notificationResolver{r} }

func (r *Resolver) Enterprise() generated.EnterpriseResolver { return &enterpriseResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type subscriptionResolver struct{ *Resolver }
type paiementResolver struct{ *Resolver }
type notificationResolver struct{ *Resolver }
type enterpriseResolver struct{ *Resolver }
