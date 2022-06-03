<template>
	<div>
		Balance: {{ balance }}
	</div>
</template>
<script lang="ts">
import { defineComponent } from "vue";
export default defineComponent({
	data() {
		return {
			balance: 1000,
		}
	},
})
</script>