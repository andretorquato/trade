import axios from "axios";

export class CoinMarketCapRemote {
	async getMarketData() {
		const response = await axios.get(
			`${import.meta.env['VITE_COIN_MARKET_CAP_URL']}/cryptocurrency/listings/latest`
		);
		if (response.status === 200) return response.data;

		return null;
	}
}