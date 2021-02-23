<template>
  <v-toolbar dense>
  
        <v-toolbar-title>Event Manager</v-toolbar-title>
        <v-spacer></v-spacer>
        <v-text-field
                v-model="filter"
                label="Filter"
                append-icon="mdi-filter"
                required
        ></v-text-field>
        
        <v-spacer></v-spacer>
  
        <v-btn icon @click="refreshAlertGroups">
            <v-icon>mdi-reload</v-icon>
        </v-btn>
      </v-toolbar>
</template>

<script>
export default {
  name: "AlertNav",
  data () {
    return {
      filter: "{}",
    }
  },
  methods: {
    refreshAlertGroups() {
        if (this.filter === "") {
            this.filter = "{}";
        }
        this.$store.dispatch('loadAlertGroups', this.filter);
    }
  },
  mounted () {
    this.$store.dispatch('loadAlertGroups', this.filter);
  }
};
</script>

