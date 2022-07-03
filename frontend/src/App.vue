<script lang="ts">
import { defineComponent } from "vue";
import Wallet from "./components/wallet/wallet.vue";
import { CoinMarketCapRemote } from "./remote/coinmarketcap-remote";
import { TradeApiRemote } from "./remote/tradeapi-remote";

export default defineComponent({
  data() {
    const marketData: any = null;
    const coinMarketCapRemote = new CoinMarketCapRemote();
    const tradeApiRemote = new TradeApiRemote();
    return { marketData, coinMarketCapRemote, tradeApiRemote };
  },
  components: {
    Wallet
  },
  methods: {
    async getMarketData() {
      const savedData = localStorage.getItem("marketData");
      if (!savedData) {
        this.coinMarketCapRemote.getMarketData().then(async (data) => {
          if (data.hasOwnProperty("error")) return 
          localStorage.setItem("marketData", JSON.stringify(data));
          
          this.marketData = data;
          const response = await this.tradeApiRemote.saveDataMarket(this.marketData);
          console.log(response);
        });
      } else {
        this.marketData = JSON.parse(savedData);
      }
    },
  },
  mounted() {
    this.getMarketData();
  },
});
</script>

<template>
  <div>
    <Wallet />
    <div v-if="marketData">
      <div v-for="coin of marketData.data" :key="coin.id">
        <p class="coin-name">
          {{ coin.name }} - price: $ {{ coin.quote.USD.price.toFixed(3) }}
        </p>
      </div>
    </div>
    <div v-else>error load</div>
  </div>
</template>

<style>
body {
  background: #181818;
  color: #fff;
}
.poc {
  display: flex;
  justify-content: center;
  font-size: 2rem;
}
</style>
