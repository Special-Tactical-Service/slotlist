<template>
	<div class="menu">
		<div class="icon"></div>
		<input type="text" v-model="title" />
		<button>{{$t("button_save")}}</button>
		<button>{{$t("button_load")}}</button>
		<button>{{$t("button_reset")}}</button>
		<button>{{$t("button_login")}}</button>
	</div>
</template>

<script>
export default {
	data() {
		return {
			title: "Slotlist"
		};
	}
}
</script>

<i18n>
{
	"en": {
		"button_save": "Save",
		"button_load": "Load",
		"button_reset": "Reset",
		"button_login": "Login"
	}
}
</i18n>
