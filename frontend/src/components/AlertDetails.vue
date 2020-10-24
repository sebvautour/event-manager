<template>
  <v-card outlined v-if="alertGroup !== null">
    <v-card-title class="d-flex justify-space-between">
      <div><h2>Alerts</h2></div>
      <div><v-btn icon @click="clearSelectedAlertGroup"><v-icon>mdi-close</v-icon></v-btn></div>
    </v-card-title>

    <v-data-table
      v-model="selected"
      :headers="headers"
      :items="alertGroup.alerts"
      item-key="id"
      show-select
      class="elevation-1"
      >

      <template v-slot:item.severity="{ item }">
        <v-chip :color="getSevColor(item.severity)" dark>{{ item.severity }}</v-chip>
      </template>

      <template v-slot:item.details="{ item }">
          {{ JSON.stringify(item.details) }}
      </template>
    </v-data-table>
  </v-card>
</template>

<script>
export default {
  name: "AlertsDetails",
  data () {
    return {
      selected: [],
      headers: [
        { text: 'Event Count', value: 'event_count' },
        { text: 'Severity', value: 'severity' },
        { text: 'Entity', value: 'entity' },
        { text: 'Message', value: 'message' },
        { text: 'Details', value: 'details' }
      ],
    }
  },
  methods: {
    getSevColor (sev) {
      switch(sev) {
        case "critical":
          return "red";
        case "major":
          return "orange";
        case "minor":
          return "amber";
        case "warning":
          return "blue";
        case "info":
          return "teal";

      }
      return "teal";
    },
    clearSelectedAlertGroup() {
      this.$store.dispatch('clearSelectedAlertGroup');
    }
  },
  computed: {
    alertGroup () {
      return this.$store.getters.getAlertGroup(this.$store.getters.getSelectedAlertGroupKey());
    }
  }
};
</script>

