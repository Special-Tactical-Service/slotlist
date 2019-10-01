<template>
    <div>
        <mainmenu></mainmenu>
        <players></players>
        <slot></slot>
        <end></end>
    </div>
</template>

<script>
import mainmenu from "./menu.vue";
import players from "./players.vue";
import end from "./end.vue";

export default {
    components: {mainmenu, players, end}
}
</script>
