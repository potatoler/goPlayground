package main

import (
	"DBwork/banking"
	"database/sql"
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

var db *sql.DB

var user banking.User

func LoginView(window fyne.Window) fyne.CanvasObject {
	title := widget.NewCard("欢迎", "账户登录", nil)
	account := widget.NewEntry()
	account.SetPlaceHolder("账号/工号")
	password := widget.NewPasswordEntry()
	password.SetPlaceHolder("密码")
	loginButton := widget.NewButton("登录", func() {
		var err error
		user, err = banking.Login(db, account.Text, password.Text)
		if err != nil {
			//TODO: Show error message
			ShowErrorMessage(window, err)
		}
		if user.Name() == "" {
			window.SetContent(LoginView(window))
		} else if user.IsAdmin() {
			window.SetContent(ClerkView(window))
		} else {
			window.SetContent(ClientView(window))
		}
	})
	return container.NewVBox(
		title,
		account,
		password,
		loginButton,
	)
}

func ClerkView(window fyne.Window) fyne.CanvasObject {
	name, account := banking.GetUserTitleStringPair(db, user)
	title := widget.NewCard(name, account, nil)
	registerButton := widget.NewButton("开户", func() {
		window.SetContent(RegisterView(window))
	})
	remittanceButton := widget.NewButton("转账", func() {
		window.SetContent(RemittanceView(window))
	})
	depositButton := widget.NewButton("存入", func() {
		window.SetContent(DepositView(window))
	})
	withdrawButton := widget.NewButton("取出", func() {
		window.SetContent(WithdrawView(window))
	})
	unblockButton := widget.NewButton("解冻", func() {
		window.SetContent(UnblockView(window))
	})
	logoutButton := widget.NewButton("退出", func() {
		user = banking.User{}
		window.SetContent(LoginView(window))
	})
	return container.NewVBox(
		title,
		registerButton,
		remittanceButton,
		depositButton,
		withdrawButton,
		unblockButton,
		logoutButton,
	)
}

func ClientView(window fyne.Window) fyne.CanvasObject {
	name, account := banking.GetUserTitleStringPair(db, user)
	title := widget.NewCard(name, account, nil)
	profileButton := widget.NewButton("账户信息", func() {
		window.SetContent(ProfileView(window))
	})
	balanceButton := widget.NewButton("财产", func() {
		window.SetContent(BalanceView(window))
	})
	logoutButton := widget.NewButton("退出", func() {
		user = banking.User{}
		window.SetContent(LoginView(window))
	})
	return container.NewVBox(
		title,
		profileButton,
		balanceButton,
		logoutButton,
	)
}

func RegisterView(window fyne.Window) fyne.CanvasObject {
	title := widget.NewCard("新账户", "", nil)
	nameEntry := widget.NewEntry()
	nameEntry.SetPlaceHolder("姓名")
	idEntry := widget.NewEntry()
	idEntry.SetPlaceHolder("证件号")
	addressEntry := widget.NewEntry()
	addressEntry.SetPlaceHolder("地址")
	telephoneEntry := widget.NewEntry()
	telephoneEntry.SetPlaceHolder("电话")
	passwordEntry := widget.NewPasswordEntry()
	passwordEntry.SetPlaceHolder("密码")
	confirmEntry := widget.NewPasswordEntry()
	confirmEntry.SetPlaceHolder("确认密码")
	nextStepButton := widget.NewButton("确认", func() {
		cardNumber, err := banking.Register(db, user,
			nameEntry.Text,
			idEntry.Text,
			addressEntry.Text,
			telephoneEntry.Text,
			passwordEntry.Text,
			confirmEntry.Text,
		)
		if err != nil {
			// TODO: Show error message
			ShowErrorMessage(window, err)
			window.SetContent(RegisterView(window))
		} else {
			window.SetContent(RegisterSuccessView(window, cardNumber))
		}
	})
	backButton := widget.NewButton("返回", func() {
		window.SetContent(ClerkView(window))
	})
	return container.NewVBox(
		title,
		nameEntry,
		idEntry,
		addressEntry,
		telephoneEntry,
		passwordEntry,
		confirmEntry,
		nextStepButton,
		backButton,
	)
}

func RegisterSuccessView(window fyne.Window, cardNumber string) fyne.CanvasObject {
	title := widget.NewCard("开户成功", "请妥善保管账号", nil)
	message := widget.NewLabel("账号：" + cardNumber)
	okButton := widget.NewButton("确认", func() {
		window.SetContent(ClerkView(window))
	})
	return container.NewVBox(
		title,
		message,
		okButton,
	)
}

func RemittanceView(window fyne.Window) fyne.CanvasObject {
	title := widget.NewCard("转账", "", nil)
	srcEntry := widget.NewEntry()
	srcEntry.SetPlaceHolder("转账账户")
	dstEntry := widget.NewEntry()
	dstEntry.SetPlaceHolder("接收账户")
	amountEntry := widget.NewEntry()
	amountEntry.SetPlaceHolder("金额")
	passwordEntry := widget.NewPasswordEntry()
	passwordEntry.SetPlaceHolder("密码")
	nextStepButton := widget.NewButton("确定", func() {
		err := banking.Remittance(db, user,
			srcEntry.Text,
			dstEntry.Text,
			amountEntry.Text,
			passwordEntry.Text,
		)
		if err != nil {
			// TODO: Show error message
			ShowErrorMessage(window, err)
			window.SetContent(RemittanceView(window))
		} else {
			// TODO: Tell success
			message := "转账成功"
			window.SetContent(SuccessMessageView(window, message))
		}
	})
	backButton := widget.NewButton("取消", func() {
		window.SetContent(ClerkView(window))
	})
	return container.NewVBox(
		title,
		srcEntry,
		dstEntry,
		amountEntry,
		passwordEntry,
		nextStepButton,
		backButton,
	)
}

func UnblockView(window fyne.Window) fyne.CanvasObject {
	title := widget.NewCard("解冻", "", nil)
	accountEntry := widget.NewEntry()
	accountEntry.SetPlaceHolder("账户")
	passwordEntry := widget.NewPasswordEntry()
	passwordEntry.SetPlaceHolder("密码")
	idEntry := widget.NewEntry()
	idEntry.SetPlaceHolder("证件")
	okButton := widget.NewButton("确定", func() {
		err := banking.Unblock(db, accountEntry.Text, passwordEntry.Text, idEntry.Text)
		if err != nil {
			// TODO: Show error message
			ShowErrorMessage(window, err)
		} else {
			window.SetContent(ClerkView(window))
		}
	})
	backButton := widget.NewButton("取消", func() {
		window.SetContent(ClerkView(window))
	})
	return container.NewVBox(
		title,
		accountEntry,
		passwordEntry,
		idEntry,
		okButton,
		backButton,
	)
}

func SuccessMessageView(window fyne.Window, message string) fyne.CanvasObject {
	title := widget.NewCard("成功", message, nil)
	okButton := widget.NewButton("确定", func() {
		window.SetContent(ClerkView(window))
	})
	return container.NewVBox(
		title,
		okButton,
	)
}

func DepositView(window fyne.Window) fyne.CanvasObject {
	title := widget.NewCard("存款", "", nil)
	accountEntry := widget.NewEntry()
	accountEntry.SetPlaceHolder("账户")
	amountEntry := widget.NewEntry()
	amountEntry.SetPlaceHolder("金额")
	passwordEntry := widget.NewPasswordEntry()
	passwordEntry.SetPlaceHolder("密码")
	okButton := widget.NewButton("确定", func() {
		err := banking.Deposit(db, user,
			accountEntry.Text,
			amountEntry.Text,
			passwordEntry.Text,
		)
		if err != nil {
			// TODO: Show error message
			ShowErrorMessage(window, err)
			window.SetContent(DepositView(window))
		} else {
			// TODO: Tell success
			message := "存入成功"
			window.SetContent(SuccessMessageView(window, message))
		}
	})
	backButton := widget.NewButton("取消", func() {
		window.SetContent(ClerkView(window))
	})
	return container.NewVBox(
		title,
		accountEntry,
		amountEntry,
		passwordEntry,
		okButton,
		backButton,
	)
}

func WithdrawView(window fyne.Window) fyne.CanvasObject {
	title := widget.NewCard("取款", "", nil)
	accountEntry := widget.NewEntry()
	accountEntry.SetPlaceHolder("账户")
	amountEntry := widget.NewEntry()
	amountEntry.SetPlaceHolder("金额")
	passwordEntry := widget.NewPasswordEntry()
	passwordEntry.SetPlaceHolder("密码")
	okButton := widget.NewButton("确定", func() {
		err := banking.Withdraw(db, user,
			accountEntry.Text,
			amountEntry.Text,
			passwordEntry.Text,
		)
		if err != nil {
			// TODO: Show error message
			ShowErrorMessage(window, err)
			window.SetContent(WithdrawView(window))
		} else {
			// TODO: Tell success
			message := "取出成功"
			window.SetContent(SuccessMessageView(window, message))
		}
	})
	backButton := widget.NewButton("取消", func() {
		window.SetContent(ClerkView(window))
	})
	return container.NewVBox(
		title,
		accountEntry,
		amountEntry,
		passwordEntry,
		okButton,
		backButton,
	)
}

func ProfileView(window fyne.Window) fyne.CanvasObject {
	client, err := banking.GetClientProfile(db, user)
	if err != nil {
		// TODO: Show error message
		ShowErrorMessage(window, err)
		return ClientView(window)
	}
	title := widget.NewCard("个人信息", TailString(user.Account(), 4), nil)
	nameLable := widget.NewLabel(fmt.Sprintf("👤姓名：%s", client.Name))
	idLable := widget.NewLabel(fmt.Sprintf("🪪证件：%s", TailString(client.Id, 4)))
	addressLable := widget.NewLabel(fmt.Sprintf("🏠地址：%s", client.Address))
	telephoneLable := widget.NewLabel(fmt.Sprintf("📱电话：%s", client.Telephone))
	editButton := widget.NewButton("编辑", func() {
		window.SetContent(ChangeProfileView(window, client))
	})
	backButton := widget.NewButton("返回", func() {
		window.SetContent(ClientView(window))
	})
	return container.NewVBox(
		title,
		nameLable,
		idLable,
		addressLable,
		telephoneLable,
		editButton,
		backButton,
	)
}

func BalanceView(window fyne.Window) fyne.CanvasObject {
	title := widget.NewCard("余额", " ", nil)
	balance, err := banking.Balance(db, user)
	if err != nil {
		// TODO: Show error message
		ShowErrorMessage(window, err)
		return ClientView(window)
	}
	balanceCard := widget.NewCard(fmt.Sprintf("%.2f", balance), fmt.Sprintf("账户：%s", TailString(user.Account(), 4)), nil)
	backButton := widget.NewButton("返回", func() {
		window.SetContent(ClientView(window))
	})
	return container.NewVBox(
		title,
		balanceCard,
		backButton,
	)
}

func ChangeProfileView(window fyne.Window, client banking.Client) fyne.CanvasObject {
	title := widget.NewCard("个人信息更改", TailString(user.Account(), 4), nil)
	nameLable := widget.NewLabel(fmt.Sprintf("👤姓名：%s", client.Name))
	idLable := widget.NewLabel(fmt.Sprintf("🪪证件：%s", TailString(client.Id, 4)))
	addressLable := widget.NewLabel("🏠地址")
	addressEntry := widget.NewEntry()
	addressEntry.SetText(client.Address)
	telephoneLable := widget.NewLabel("📱电话")
	telephoneEntry := widget.NewEntry()
	telephoneEntry.SetText(client.Telephone)
	okButton := widget.NewButton("确定", func() {
		err := banking.ChangeProfile(db, client.Id, telephoneEntry.Text, addressEntry.Text)
		if err != nil {
			// TODO: Show error message
			ShowErrorMessage(window, err)
		} else {
			window.SetContent(ProfileView(window))
		}
	})
	backButton := widget.NewButton("取消", func() {
		window.SetContent(ProfileView(window))
	})
	return container.NewVBox(
		title,
		nameLable,
		idLable,
		addressLable,
		addressEntry,
		telephoneLable,
		telephoneEntry,
		okButton,
		backButton,
	)
}

func ConnectingView(window fyne.Window) fyne.CanvasObject {
	title := widget.NewLabel("connecting")
	connString := "server=localhost;user id=SA;password=Ren20040525;port=1433;database=bank"
	var err error
	db, err = sql.Open("mssql", connString)
	if err != nil {
		fmt.Printf("Failed connecting to the database\n")
		panic(err.Error())
	}
	fmt.Printf("conncted.")
	return container.NewVBox(
		title,
	)
}
