<template>
  <div>
    <v-data-table
      v-model="selected"
      :headers="headers"
      :items="alertGroups"
      item-key="group_key"
      show-select
      class="elevation-1"
      >

        <template v-slot:item.primary_alert.severity="{ item }">
          <v-chip :color="getSevColor(item.primary_alert.severity)" dark>{{ item.primary_alert.severity }}</v-chip>
        </template>

        <template v-slot:item.action="{ item }">
          <v-icon
          small
          class="mr-2"
          @click="selectAlertGroup(item)"
          >
            mdi-magnify
          </v-icon>
        </template>
    </v-data-table>

  </div>
</template>

<script>
export default {
  name: "AlertsTable",
  props: {},
  data () {
    return {
      selected: [],
      headers: [
        { text: 'Alert Count', value: 'alert_count' },
        { text: 'Severity', value: 'primary_alert.severity' },
        { text: 'Entity', value: 'primary_alert.labels.entity' },
        { text: 'Message', value: 'primary_alert.labels.message' },
        { text: '', value: 'action', sortable: false },
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
    selectAlertGroup(ag) {
      this.$store.dispatch('setSelectedAlertGroup', ag.group_key);
    }
  },
  computed: {
    alertGroups () {
      return this.$store.getters.getAlertGroups();
    }
  },
};
</script>

