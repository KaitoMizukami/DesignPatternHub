""" 

問題: 銀行アプリの口座管理システムの実装

あなたは、銀行アプリの口座管理システムをPythonで実装する必要があります。
このシステムは、顧客の銀行口座を管理し、預金や引き出し、残高照会などの機能を提供します。

口座データは、顧客の情報と残高を持っています：

口座番号
顧客名
残高
また、預金と引き出しのトランザクションは履歴として保存します。

このシステムの要件は以下の通りです：

口座の新規作成ができる。
口座に預金や引き出しを行える。
口座の残高を照会できる。
口座のトランザクション履歴を参照できる。
口座の種類（普通預金、定期預金など）を追加する際に、拡張性を考慮した実装を行う。

上記の要件を満たすために、デザインパターンを使用してシステムを実装してください。

"""
from abc import ABC, abstractmethod


# 口座の種類
class AccountType:
    SAVINGS = "Savings"
    CHECKING = "Checking"
    FIXED_DEPOSIT = "Fixed Deposit"

# 口座データのクラス
class Account:
    def __init__(self, account_number, customer_name, balance, account_type):
        self.account_number = account_number
        self.customer_name = customer_name
        self.balance = balance
        self.account_type = account_type
        self.transactions = []

    def deposit(self, amount):
        self.balance += amount
        self.transactions.append(f"Deposit: {amount}")

    def withdraw(self, amount):
        if self.balance >= amount:
            self.balance -= amount
            self.transactions.append(f"Withdraw: {amount}")
        else:
            print("Insufficient balance.")

    def get_balance(self):
        return self.balance

    def get_transaction_history(self):
        return self.transactions

# 口座作成のファクトリメソッド
class AccountFactory:
    @staticmethod
    def create_account(account_number, customer_name, balance, account_type):
        return Account(account_number, customer_name, balance, account_type)

# 預金コマンド
class DepositCommand:
    def __init__(self, account, amount):
        self.account = account
        self.amount = amount

    def execute(self):
        self.account.deposit(self.amount)

# 引き出しコマンド
class WithdrawCommand:
    def __init__(self, account, amount):
        self.account = account
        self.amount = amount

    def execute(self):
        self.account.withdraw(self.amount)

# 取引履歴イテレータ
class TransactionIterator:
    def __init__(self, account):
        self.account = account
        self.index = 0

    def __iter__(self):
        return self

    def __next__(self):
        if self.index < len(self.account.get_transaction_history()):
            transaction = self.account.get_transaction_history()[self.index]
            self.index += 1
            return transaction
        else:
            raise StopIteration

if __name__ == "__main__":
    # 口座作成
    account1 = AccountFactory.create_account("123456789", "Alice", 10000, AccountType.SAVINGS)
    account2 = AccountFactory.create_account("987654321", "Bob", 5000, AccountType.CHECKING)

    # 預金と引き出し
    deposit_command1 = DepositCommand(account1, 2000)
    deposit_command2 = DepositCommand(account2, 3000)
    withdraw_command1 = WithdrawCommand(account1, 500)
    withdraw_command2 = WithdrawCommand(account2, 1000)

    deposit_command1.execute()
    deposit_command2.execute()
    withdraw_command1.execute()
    withdraw_command2.execute()

    # 口座残高と取引履歴の表示
    print(f"{account1.customer_name}の{account1.account_type}の残高: {account1.get_balance()}円")
    print(f"{account2.customer_name}の{account2.account_type}の残高: {account2.get_balance()}円")

    print(f"{account1.customer_name}の取引履歴:")
    for transaction in TransactionIterator(account1):
        print(transaction)

    print(f"{account2.customer_name}の取引履歴:")
    for transaction in TransactionIterator(account2):
        print(transaction)
