package executor

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/kupriyanovkk/gophkeeper/internal/client/app"
	"github.com/kupriyanovkk/gophkeeper/internal/client/model"
	"github.com/kupriyanovkk/gophkeeper/internal/client/storage"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Executor struct {
	app *app.App
}

// Execute executes the command given as a string.
//
// s string
func (e *Executor) Execute(s string) {
	var isForce bool
	cmd, args := parseCommandArgs(s)

	if args["force"] || args["f"] {
		isForce = true
	}

	switch cmd[0] {
	case "login":
		if err := e.login(cmd); err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println("login success")

	case "register":
		if err := e.register(cmd); err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println("register success")

	case "create-login-pass":
		if err := e.createLoginPass(cmd); err != nil {
			fmt.Println(err)
			return
		}

	case "create-card":
		if err := e.createCard(cmd); err != nil {
			fmt.Println(err)
			return
		}

	case "create-text":
		if err := e.createText(cmd); err != nil {
			fmt.Println(err)
			return
		}

	case "create-file":
		if err := e.createFile(cmd); err != nil {
			fmt.Println(err)
			return
		}

	case "get-private":
		private, err := e.getPrivate(cmd)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Printf("Private data:%+v\n", private)

	case "get-private-by-type":
		list, err := e.getPrivateByType(cmd)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Printf("Private data list:\n")
		for _, private := range list {
			fmt.Printf("ID:%v Title: %v\n", private.Id, private.Title)
		}

	case "get-private-binary":
		if err := e.getPrivateBinary(cmd); err != nil {
			fmt.Println(err)
			return
		}

	case "update-private":
		if err := e.updatePrivate(cmd, isForce); err != nil {
			fmt.Println(err)
			return
		}

	case "delete-private":
		if err := e.deletePrivate(cmd); err != nil {
			fmt.Println(err)
			return
		}

	case "exit":
		fmt.Println("goodbye, see you soon")

		e.app.Cancel()
		e.app.Cron.Stop()

		os.Exit(0)
	}
}

// login performs user login.
//
// args []string - the arguments for login.
// error - returns an error if the login fails.
func (e *Executor) login(args []string) error {
	if len(args)-1 < 2 {
		return fmt.Errorf("login: missing required arguments")
	}

	user := model.User{Login: args[1], Password: args[2]}

	if err := e.app.UserService.Login(user); err != nil {
		st, _ := status.FromError(err)

		if st.Code() == codes.NotFound {
			return fmt.Errorf("login: user not found")
		}

		return fmt.Errorf("login: error %w", err)
	}

	e.app.Sync.SyncAll()

	go e.app.Cron.Run()

	return nil
}

// register registers a user with the provided arguments.
//
// args []string - the arguments for user registration.
// error - returns an error if registration fails.
func (e *Executor) register(args []string) error {
	if len(args)-1 < 2 {
		return fmt.Errorf("register: missing required arguments")
	}

	user := model.User{Login: args[1], Password: args[2]}

	if err := e.app.UserService.Register(user); err != nil {
		if status.Code(err) == codes.InvalidArgument {
			return fmt.Errorf("register: invalid arguments %w", err)
		}

		return err
	}

	return nil
}

// createLoginPass creates a login pass for the Executor.
//
// It takes a slice of strings 'args' as parameter and returns an error.
func (e *Executor) createLoginPass(args []string) error {
	if len(args)-1 < 3 {
		return fmt.Errorf("create-login-pass: missing required arguments")
	}

	loginPass := model.PrivateLoginPass{
		Title:    args[1],
		Login:    args[2],
		Password: args[3],
		Type:     int(storage.PrivateLoginPass),
	}

	content, errMarshal := json.Marshal(loginPass)
	if errMarshal != nil {
		return errMarshal
	}

	return e.app.PrivateService.CreatePrivate(loginPass.Title, loginPass.Type, string(content))
}

// createCard creates a private card with the given arguments.
//
// args []string
// error
func (e *Executor) createCard(args []string) error {
	if len(args)-1 < 4 {
		return fmt.Errorf("create-card: missing required arguments")
	}

	card := model.PrivateCard{
		Title:      args[1],
		CardNumber: args[2],
		CVV:        args[3],
		Due:        args[4],
		Type:       int(storage.PrivateCard),
	}

	content, errMarshal := json.Marshal(card)
	if errMarshal != nil {
		return errMarshal
	}

	return e.app.PrivateService.CreatePrivate(card.Title, card.Type, string(content))
}

// createText creates a text with the given arguments.
//
// Parameter(s): args []string
// Return type: error
func (e *Executor) createText(args []string) error {
	if len(args)-1 < 2 {
		return fmt.Errorf("create-text: missing required arguments")
	}

	text := model.PrivateText{
		Title: args[1],
		Text:  strings.Join(args[2:], " "),
		Type:  int(storage.PrivateText),
	}

	content, errMarshal := json.Marshal(text)
	if errMarshal != nil {
		return errMarshal
	}

	return e.app.PrivateService.CreatePrivate(text.Title, text.Type, string(content))
}

// createFile creates a file based on the provided arguments.
//
// args []string
// error
func (e *Executor) createFile(args []string) error {
	if len(args) < 2 {
		return fmt.Errorf("create-file: missing required arguments")
	}

	file := model.PrivateFile{
		Title: args[1],
		Path:  args[2],
		Type:  int(storage.PrivateFile),
	}

	data, err := os.ReadFile(args[2])
	if err != nil {
		return err
	}

	return e.app.PrivateService.CreatePrivate(file.Title, file.Type, string(data))
}

// getPrivateBinary is a function to retrieve a private binary.
//
// It takes an array of strings as the arguments and returns an error.
func (e *Executor) getPrivateBinary(args []string) error {
	if len(args) < 2 {
		return fmt.Errorf("get-private-binary: missing required arguments")
	}

	id, parseErr := strconv.Atoi(args[1])
	if parseErr != nil {
		return parseErr
	}

	return e.app.PrivateService.GetPrivateBinary(id, args[2])
}

// getPrivate retrieves private data.
//
// args []string
// (interface{}, error)
func (e *Executor) getPrivate(args []string) (interface{}, error) {
	if len(args) < 2 {
		return nil, fmt.Errorf("get-private: missing required arguments")
	}

	id, err := strconv.Atoi(args[1])
	if err != nil {
		return nil, err
	}

	private, err := e.app.PrivateService.GetPrivateData(id)
	if err != nil {
		st, _ := status.FromError(err)
		if st.Code() == codes.NotFound {
			return nil, fmt.Errorf("record not found")
		}

		return nil, err
	}

	return private, nil
}

// getPrivateByType retrieves private data based on the provided type.
//
// args: a slice of strings representing the type and any additional arguments.
// []model.PrivateDataList, error: returns a slice of private data and an error if any.
func (e *Executor) getPrivateByType(args []string) ([]model.PrivateDataList, error) {
	if len(args) < 1 {
		return nil, fmt.Errorf("get-private-by-type: missing required arguments")
	}

	id, err := strconv.Atoi(args[1])
	if err != nil {
		return nil, err
	}

	data, err := e.app.PrivateService.GetPrivateDataList(id)
	if err != nil {
		return nil, err
	}

	var list []model.PrivateDataList
	for _, d := range data {
		list = append(list, model.PrivateDataList{
			Id:    int(d.Id),
			Title: d.Title,
		})
	}

	return list, nil
}

// updatePrivate edits a private entity based on the provided arguments and isForce flag.
//
// Parameters:
// - args []string: the arguments for editing the private entity
// - isForce bool: a flag indicating whether to force the edit operation
// Return type: error
func (e *Executor) updatePrivate(args []string, isForce bool) error {
	var (
		privateType int
		privateID   int
		body        []byte
	)

	numArgs := len(args) - 1
	if numArgs >= 3 {
		privateType, err := strconv.Atoi(args[3])
		if err != nil {
			return err
		}

		privateID, err = strconv.Atoi(args[1])
		if err != nil {
			return err
		}

		switch privateType {
		case int(storage.PrivateLoginPass):
			switch numArgs {
			case 4:
				return fmt.Errorf("validation error: Password is missing")

			case 3:
				return fmt.Errorf("validation error: Login and Password is missing")

			default:
				p := model.PrivateLoginPass{
					Id:       privateID,
					Title:    args[2],
					Type:     privateType,
					Login:    args[4],
					Password: args[5],
				}
				body, err = json.Marshal(p)
				if err != nil {
					return err
				}
			}

		case int(storage.PrivateText):
			switch numArgs {
			case 3:
				return fmt.Errorf("validation error: Text is missing")

			default:
				p := model.PrivateText{
					Id:    privateID,
					Title: args[2],
					Type:  privateType,
					Text:  strings.Join(args[4:], " "),
				}
				body, err = json.Marshal(p)
				if err != nil {
					return err
				}
			}

		case int(storage.PrivateFile):
			switch numArgs {
			case 3:
				return fmt.Errorf("validation error: Filepath is missing")

			default:
				p := model.PrivateFile{
					Id:    privateID,
					Title: args[2],
					Type:  privateType,
					Path:  args[3],
				}
				body, err = json.Marshal(p)
				if err != nil {
					return err
				}
			}

		case int(storage.PrivateCard):
			switch numArgs {
			case 5:
				return fmt.Errorf("validation error: Due date is missing")

			case 4:
				return fmt.Errorf("validation error: CVV and Due date is missing")

			case 3:
				return fmt.Errorf("validation error: Card number, CVV and Due date is missing")

			default:
				p := model.PrivateCard{
					Id:         privateID,
					Title:      args[2],
					Type:       privateType,
					CardNumber: args[4],
					CVV:        args[5],
					Due:        args[6],
				}
				body, err = json.Marshal(p)
				if err != nil {
					return err
				}
			}
		}
	} else {
		return fmt.Errorf("validation error: fields is missing")
	}

	if err := e.app.PrivateService.UpdatePrivateData(privateID, args[3], privateType, string(body), isForce); err != nil {
		st, _ := status.FromError(err)
		fmt.Println(st.Message())

		if st.Code() == codes.FailedPrecondition {
			fmt.Println("starting re-sync")
			e.app.Sync.SyncAll()
			fmt.Println("re-sync ended")
		}

		return nil
	}

	return nil
}

// deletePrivate deletes a private item.
//
// args []string
// error
func (e *Executor) deletePrivate(args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("delete-private: missing required arguments")
	}

	id, convErr := strconv.Atoi(args[1])
	if convErr != nil {
		return convErr
	}

	return e.app.PrivateService.DeletePrivate(id)
}

// NewExecutor creates a new Executor.
//
// No parameters.
// Returns a pointer to Executor.
func NewExecutor() *Executor {
	clientApp, err := app.NewApp()
	if err != nil {
		panic(err)
	}

	return &Executor{app: clientApp}
}

func parseCommandArgs(cmd string) ([]string, map[string]bool) {
	args := make(map[string]bool)
	cmds := strings.Fields(strings.TrimSpace(cmd))
	filtered := make([]string, 0, len(cmds))

	for _, c := range cmds {
		if strings.HasPrefix(c, "-") {
			opt := strings.TrimPrefix(c, "-")
			values := strings.SplitN(opt, "=", 2)
			args[values[0]] = true
			continue
		}

		filtered = append(filtered, c)
	}

	return filtered, args
}
