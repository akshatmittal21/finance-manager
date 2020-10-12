<template>
  <div>
    <div class="container">
      <b-tabs fill v-model="currentIndex">
        <b-tab title="Transactions" active>
          <p>
            <Transactions v-if="currentIndex == 0" />
          </p>
        </b-tab>
        <b-tab title="Accounts">
          <p>
            <Accounts v-if="currentIndex == 1" />
          </p>
        </b-tab>
        <b-tab title="Transaction Categories">
          <p>
            <Category v-if="currentIndex == 2" />
          </p>
        </b-tab>
      </b-tabs>
    </div>
  </div>
</template>

<script>
// import axios from 'axios'
import Transactions from '@/components/Transactions.vue'
import Accounts from '@/components/Accounts.vue'
import Category from '@/components/Category.vue'
export default {
  name: 'FinanceManager',
  components: {
    Transactions,
    Accounts,
    Category
  },
  data() {
    return {
      currentIndex: 0
    }
  }
}
</script>

<style lang="scss" scoped></style>
