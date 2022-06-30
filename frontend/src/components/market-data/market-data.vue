<template>
	<div v-if="marketData">
		{{ marketData }} 
	</div>
</template>
<script lang="ts">
import { defineComponent } from "vue";
export default defineComponent({
	data() {
		const marketData: any = null;
		return { marketData }
	},
	methods: {

	}
})
</script>