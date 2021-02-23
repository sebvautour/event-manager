<template>
  <div>
    <v-main>
      <AlertNav/>

      <v-row
        no-gutters
        style="flex-wrap: nowrap;"
      >
        <v-col
          cols="12"
          style="min-width: 100px;"
          class="flex-grow-0 flex-shrink-1"
        >
       <AlertsTable/>
        </v-col>
       
        <v-col
          class="flex-grow-0 flex-shrink-0"
        >
        <AlertDetails/>
        </v-col>
      </v-row>

    </v-main>

  </div>
</template>

<script>
// @ is an alias to /src
import AlertsTable from "@/components/AlertsTable.vue";
import AlertNav from "@/components/AlertNav.vue";

import AlertDetails from "@/components/AlertDetails.vue";

export default {
  name: "Console",
  components: {
    AlertsTable,
    AlertDetails,
    AlertNav
  }
};
</script>
