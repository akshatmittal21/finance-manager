<template>
  <div>
    <div>
      <b-button
        block
        variant="primary"
        size="sm"
        @click="openModal('addNewTransaction')"
      >+ Add New Transaction</b-button>
    </div>
    <br />

    <b-list-group>
      <b-list-group-item
        v-for="(transaction, index) in transactions"
        :key="index"
        class="transaction"
      >
        <div class="row">
          <div class="col-md-11" @dblclick="updateModal(transaction)">
            <div class="d-flex justify-content-between align-items-center">
              <strong>{{ transaction.transactionName }}</strong>
              <b-badge
                :variant="
                  transaction.transactionType === 'TRANSFER'
                    ? 'primary'
                    : transaction.transactionType === 'EXPENSE'
                    ? 'danger'
                    : 'success'
                "
                pill
              >{{ "₹ " + transaction.amount }}</b-badge>
            </div>
            <div class="small text-muted d-flex justify-content-between align-items-center">
              <strong>{{ transaction.categoryInfo.categoryName }}</strong>

              <div>
                {{ transaction.accountInfo.accountName }}
                <b-badge pill>
                  {{
                  "[₹ " + remainingBalance(index) + "]"
                  }}
                </b-badge>
              </div>
            </div>
          </div>
          <div class="col-md-1">
            <div class="row flex-row-reverse">
              <svg
                @click="deleteTransaction(transaction)"
                class="bi bi-trash-fill delete mr-2"
                width="1em"
                height="1em"
                viewBox="0 0 16 16"
                fill="currentColor"
                xmlns="http://www.w3.org/2000/svg"
              >
                <path
                  fill-rule="evenodd"
                  d="M2.5 1a1 1 0 00-1 1v1a1 1 0 001 1H3v9a2 2 0 002 2h6a2 2 0 002-2V4h.5a1 1 0 001-1V2a1 1 0 00-1-1H10a1 1 0 00-1-1H7a1 1 0 00-1 1H2.5zm3 4a.5.5 0 01.5.5v7a.5.5 0 01-1 0v-7a.5.5 0 01.5-.5zM8 5a.5.5 0 01.5.5v7a.5.5 0 01-1 0v-7A.5.5 0 018 5zm3 .5a.5.5 0 00-1 0v7a.5.5 0 001 0v-7z"
                  clip-rule="evenodd"
                />
              </svg>

              <!-- <a  download="_blank">jhkj</a> -->
              <svg
                @click="downloadReceipt(transaction.receiptPath)"
                v-if="transaction.receiptPath !== '' && transaction.receiptPath"
                class="bi bi-download mr-3 download"
                width="1em"
                height="1em"
                viewBox="0 0 16 16"
                fill="currentColor"
                xmlns="http://www.w3.org/2000/svg"
              >
                <path
                  fill-rule="evenodd"
                  d="M.5 8a.5.5 0 01.5.5V12a1 1 0 001 1h12a1 1 0 001-1V8.5a.5.5 0 011 0V12a2 2 0 01-2 2H2a2 2 0 01-2-2V8.5A.5.5 0 01.5 8z"
                  clip-rule="evenodd"
                />
                <path
                  fill-rule="evenodd"
                  d="M5 7.5a.5.5 0 01.707 0L8 9.793 10.293 7.5a.5.5 0 11.707.707l-2.646 2.647a.5.5 0 01-.708 0L5 8.207A.5.5 0 015 7.5z"
                  clip-rule="evenodd"
                />
                <path
                  fill-rule="evenodd"
                  d="M8 1a.5.5 0 01.5.5v8a.5.5 0 01-1 0v-8A.5.5 0 018 1z"
                  clip-rule="evenodd"
                />
              </svg>
            </div>
          </div>
        </div>
      </b-list-group-item>
    </b-list-group>
    <b-modal
      id="addNewTransaction"
      ref="addNewTransaction"
      class="modal-backdrop"
      centered
      title="Add New Transaction"
      modal-class="custom-modal"
      dialog-class="personalInfo-modal"
    >
      <!-- ok-title="Confirm"
      @ok.prevent="UpdateApplicantInformation()"-->
      <form>
        <div class="card-body">
          <div class="row mb-2">
            <div class="form-group col-md-12 col-lg-12 mb-md-2 mb-lg-0">
              <label class="input-label">Transaction Name</label>
              <!-- <div class="summary-data">{{schoolArea || 'Not Available'}}</div> -->
              <input
                class="form-control"
                type="text"
                name="schoolArea"
                v-model="newTransactionDetails.transactionName"
              />
            </div>

            <div class="form-group col-md-12 col-lg-12 mt-2 mt-md-0">
              <label class="input-label">Amount</label>
              <!-- <div class="summary-data">{{applicantStatus || 'Not Availabe'}}</div> -->
              <input class="form-control" type="number" v-model="newTransactionDetails.amount" />
            </div>
            <div class="form-group col-md-12 col-lg-12 mt-2 mt-md-0">
              <label class="input-label">Category</label>
              <vSelect
                label="categoryName"
                :options="categoryList"
                v-model="category"
                placeholder="-- Select Category --"
                name="category"
              />
            </div>
            <div class="form-group col-md-12 col-lg-12 mt-2 mt-md-0">
              <label class="input-label">Account</label>
              <vSelect
                label="accountName"
                :options="accountList"
                v-model="account"
                placeholder="-- Select Account --"
                name="account"
              />
            </div>
            <div class="form-group col-md-12 col-lg-12 mt-2 mt-md-0">
              <label class="input-label">Type</label>
              <vSelect
                :options="transactionTypeList"
                v-model="transactionType"
                placeholder="-- Select Type --"
                name="type"
              />
            </div>
            <div
              class="form-group col-md-12 col-lg-12 mt-2 mt-md-0"
              v-if="transactionType === 'TRANSFER'"
            >
              <label class="input-label">Transfer Amount From Account</label>
              <vSelect
                label="accountName"
                :options="accountList"
                v-model="transferAccount"
                placeholder="-- Select Account --"
                name="account"
              />
            </div>
            <div class="form-group col-md-12 col-lg-12 mt-2 mt-md-0">
              <label class="input-label">Upload Receipt</label>
              <input class="form-control" type="file" ref="file" id="file" name="filename" />
            </div>
          </div>
        </div>
      </form>
      <template v-slot:modal-footer="{ ok, cancel }">
        <b-button variant="success" @click="modelData()">Add</b-button>
        <b-button variant="secondary" @click="cancel()">Cancel</b-button>
      </template>
    </b-modal>
    <b-modal
      id="updateTransaction"
      ref="updateTransaction"
      class="modal-backdrop"
      centered
      title="Update Transaction"
      modal-class="custom-modal"
      dialog-class="personalInfo-modal"
    >
      <!-- ok-title="Confirm"
      @ok.prevent="UpdateApplicantInformation()"-->
      <form>
        <div class="card-body">
          <div class="row mb-2">
            <div class="form-group col-md-12 col-lg-12 mb-md-2 mb-lg-0">
              <label class="input-label">Transaction Name</label>
              <!-- <div class="summary-data">{{schoolArea || 'Not Available'}}</div> -->
              <input
                class="form-control"
                type="text"
                name="schoolArea"
                v-model="updateTransaction.transactionName"
              />
            </div>

            <div class="form-group col-md-12 col-lg-12 mt-2 mt-md-0">
              <label class="input-label">Amount</label>
              <!-- <div class="summary-data">{{applicantStatus || 'Not Availabe'}}</div> -->
              <input class="form-control" type="number" v-model="updateTransaction.amount" />
            </div>
            <div class="form-group col-md-12 col-lg-12 mt-2 mt-md-0">
              <label class="input-label">Category</label>
              <vSelect
                label="categoryName"
                :options="categoryList"
                v-model="updateTransaction.categoryInfo"
                placeholder="-- Select Category --"
                name="category"
              />
            </div>
            <div class="form-group col-md-12 col-lg-12 mt-2 mt-md-0">
              <label class="input-label">Account</label>
              <vSelect
                label="accountName"
                :options="accountList"
                v-model="updateTransaction.accountInfo"
                placeholder="-- Select Account --"
                name="account"
              />
            </div>
            <div class="form-group col-md-12 col-lg-12 mt-2 mt-md-0">
              <label class="input-label">Type</label>
              <vSelect
                :options="transactionTypeList"
                v-model="updateTransaction.transactionType"
                placeholder="-- Select Type --"
                name="type"
              />
            </div>
            <div
              class="form-group col-md-12 col-lg-12 mt-2 mt-md-0"
              v-if="transactionType === 'TRANSFER'"
            >
              <label class="input-label">Transfer Amount From Account</label>
              <vSelect
                label="accountName"
                :options="accountList"
                v-model="updateTransaction.transferredFromAccountInfo"
                placeholder="-- Select Account --"
                name="account"
              />
            </div>
          </div>
        </div>
      </form>
      <template v-slot:modal-footer="{ ok, cancel }">
        <b-button variant="success" @click="updateData()">Update</b-button>
        <b-button variant="secondary" @click="cancel()">Cancel</b-button>
      </template>
    </b-modal>
  </div>
</template>

<script>
import vSelect from 'vue-select'

export default {
  name: 'Transactions',
  data() {
    return {
      transactions: [],
      newTransactionDetails: {},
      accountList: [],
      categoryList: [],
      account: null,
      category: null,
      transactionType: null,
      transferAccount: null,
      transactionTypeList: ['EXPENSE', 'INCOME', 'TRANSFER'],
      updateTransaction: {}
    }
  },
  components: {
    // EditPersonalInformation,
    vSelect
  },

  methods: {
    remainingBalance(index) {
      let oldT = this.transactions
        .slice(index, this.transactions.length)
        .filter(
          a =>
            a.accountInfo.accountId ===
            this.transactions[index].accountInfo.accountId
        )

      let income = oldT
        .filter(a => a.isAmountReceived)
        .map(a => parseFloat(a.amount))

      let expense = oldT
        .filter(a => !a.isAmountReceived)
        .map(a => parseFloat(a.amount))

      income = income.length
        ? income.reduce((acc, val) => {
            return acc + val
          })
        : 0
      expense = expense.length
        ? expense.reduce((acc, val) => {
            return acc + val
          })
        : 0
      //   console.log(income, expense)
      //   let balance = this.accountList.filter(
      //     a => a.accountId === this.transactions[index].accountInfo.accountId
      //   )[0].latestBalance
      return income - expense
    },
    openModal(name) {
      this.$refs[name].show()
    },
    hideModal(name) {
      this.$refs[name].hide()
    },
    getMasterData() {
      this.axios.get('server/getAllAccounts').then(response => {
        this.accountList = response.data.map(a => ({
          accountId: a.accountId,
          accountName: a.accountInfo.accountName,
          accountType: a.accountInfo.accountType
        }))
      })
      this.axios.get('server/getAllCategories').then(response => {
        this.categoryList = response.data.map(a => ({
          categoryId: a.categoryId,
          categoryName: a.categoryInfo.categoryName
        }))
      })
    },
    modelData() {
      if (this.transactionType === 'TRANSFER') {
        let data = this.newTransactionDetails
        data.amount = parseFloat(data.amount)
        data.transactionType = this.transactionType
        data.accountInfo = this.account
        // data.transferredToAccountInfo = this.account
        data.transferredFromAccountInfo = this.transferAccount
        data.isAmountReceived = true
        data.categoryInfo = this.category

        this.addNewTransaction(data).then(res => {
          data.amount = parseFloat(data.amount)
          data.isAmountReceived = false
          data.accountInfo = this.transferAccount
          // data.transferredToAccountInfo = this.transferAccount
          data.transferredToAccountInfo = this.account
          console.log(data, res)

          this.addNewTransaction(data).then(res => {
            this.transferAccount = null || res
            this.getAllTransactions()

            this.hideModal('addNewTransaction')
            this.newTransactionDetails = {}
            this.account = null
            this.category = null
            this.transactionType = null
          })
        })
      } else {
        let data = this.newTransactionDetails
        data.amount = parseFloat(data.amount)
        data.transactionType = this.transactionType
        data.accountInfo = this.account
        // data.transferredToAccountInfo = this.account
        data.categoryInfo = this.category
        data.isAmountReceived = data.transactionType === 'INCOME'

        this.addNewTransaction(data).then(res => {
          console.log(res)
          this.getAllTransactions()

          this.hideModal('addNewTransaction')
          this.newTransactionDetails = {}
          this.account = null
          this.category = null
          this.transactionType = null
        })
      }
    },
    addNewTransaction(data) {
      let self = this
      var formData = new FormData()
      var file = this.$refs.file.files[0]
      formData.append('file', file)

      return new Promise(resolve => {
        this.axios
          .post('server/addNewTransaction', data)
          .then(function(response) {
            console.log(response.data)
            resolve()
            if (file) {
              formData.append('name', response.data)

              self.axios.post('server/uploadFile', formData, {
                headers: {
                  'Content-Type': 'multipart/form-data'
                }
              })
            }
          })
          .catch(function(error) {
            resolve()
            console.log(error)
          })
      })
    },
    getAllTransactions() {
      this.axios.get('server/getAllTransactions').then(response => {
        this.transactions = response.data || []
        this.transactions = this.transactions.sort(function(a, b) {
          return b.transactionDate - a.transactionDate
        })
      })
    },
    updateModal(transaction) {
      this.updateTransaction = transaction
      this.openModal('updateTransaction')
    },
    updateData() {
      let self = this
      this.updateTransaction.amount = parseFloat(this.updateTransaction.amount)

      return new Promise(resolve => {
        this.axios
          .post('server/updateTransaction', self.updateTransaction)
          .then(function(response) {
            console.log(response.data)
            resolve()
            self.getAllTransactions()
            self.hideModal('updateTransaction')
          })
          .catch(function(error) {
            resolve()
            console.log(error)
          })
      })
    },
    deleteTransaction(data) {
      let self = this
      return new Promise(resolve => {
        this.axios
          .post('server/deleteTransaction', data)
          .then(function(response) {
            console.log(response.data)
            resolve()
            self.getAllTransactions()
          })
          .catch(function(error) {
            resolve()
            console.log(error)
          })
      })
    },
    downloadReceipt(path) {
      // delete link
      this.axios({
        url: 'server/files/' + path,
        method: 'GET',
        responseType: 'blob'
      }).then(response => {
        let blob = new Blob([response.data])
        var fileURL = window.URL.createObjectURL(blob)

        var fileLink = document.createElement('a')

        fileLink.href = fileURL
        fileLink.setAttribute('download', path)
        document.body.appendChild(fileLink)

        fileLink.click()
      })
    }
  },
  mounted() {
    this.getMasterData()
    this.getAllTransactions()
  }
}
</script>

<style scoped>
.transaction {
  /* cursor: pointer; */
}
/* .transaction:hover {
  border: 1px solid #d0d0d0;
  padding-top: 0;
} */
.transaction:hover {
  border-color: #66afe9;
  box-shadow: 0 0 8px rgba(0, 162, 255, 0.6);
}
.delete {
  cursor: pointer;
}
.download {
  cursor: pointer;
}
</style>
