<script lang="ts">
import { defineComponent } from "vue";
import axios from "axios";

export default defineComponent({
  data() {
    const marketData: any = null;
    return { marketData };
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
  <div v-if="marketData">
    <div v-for="coin of marketData.data" :key="coin.id">
      <p class="coin-name">
        {{ coin.name }} - price: $ {{ coin.quote.USD.price.toFixed(3) }}
      </p>
    </div>
  </div>
  <div v-else> error load </div>
</template>

<style>
.poc {
  display: flex;
  justify-content: center;
  font-size: 2rem;
}
</style>
