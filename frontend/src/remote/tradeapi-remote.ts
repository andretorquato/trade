import axios from "axios";

export class TradeApiRemote {
	async saveDataMarket(data: any) {
		const response = await axios.post(`${import.meta.env['VITE_TRADE_API_URL']}/new-market-data`, { 
			data
		});
		if (response.status === 200) return response.data;

		return null;
	}
}