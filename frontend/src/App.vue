<script lang="ts">
import { defineComponent } from "vue";
import axios from "axios";
import Wallet from "./components/wallet/wallet.vue";
// import MarketCapController from "./controllers/market-cap/market-cap-controller";

export default defineComponent({
  data() {
    const marketData: any = null;
    return { marketData };
  },
  components: {
    Wallet
  },
  methods: {
    async getMarketData() {
      const savedData = localStorage.getItem("marketData");
      if (!savedData) {
        const response = await axios.get(
          "http://localhost:5000/cryptocurrency/listings/latest"
        );
        if (response.status === 200) {
          localStorage.setItem("marketData", JSON.stringify(response.data));
          this.marketData = response.data;
        }
      } else {
        // MarketCapController.getMarketCap({}, {});
        this.marketData = JSON.parse(savedData);
        console.log(this.marketData);
        // one hour = 3600000 ms
        // struct
        // {
        //
        // }
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
