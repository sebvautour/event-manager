import Vue from "vue";
import Vuex from "vuex";
import axios from 'axios';

Vue.use(Vuex);

import { baseAPIv1 } from '../variables.js'


export default new Vuex.Store({
  state: {
    alertGroups: [],
    selectedAlertGroup: null,
  },
  getters: {
    getAlertGroups: (state)  => () => {
      return state.alertGroups;
    },
    getAlertGroup: (state) => (alertGroupKey) => {
      if (!alertGroupKey) {
        return null;
      }
      return state.alertGroups.find(ag => ag.group_key === alertGroupKey);
    },
    getSelectedAlertGroupKey: (state)  => () => {
      return state.selectedAlertGroup;
    },
  },
  mutations: {
    SET_ALERT_GROUPS (state, alertsGrouped) {
      state.alertGroups = alertsGrouped["groups"];
    },
    SET_SELECTED_ALERT_GROUP (state, alertGroup) {
      state.selectedAlertGroup = alertGroup;
    },
  },
  actions: {
    loadAlertGroups ({ commit }) {
      axios
      .get(baseAPIv1+'/alerts?filter={}&grouped=true')
      .then(response => {
        commit('SET_ALERT_GROUPS', response.data);
      })
      .catch(error => {
        console.log(error);
      })
    },
    setSelectedAlertGroup ({ commit }, alertGroupKey) {
      commit('SET_SELECTED_ALERT_GROUP', alertGroupKey);
    },
    clearSelectedAlertGroup ({ commit }) {
      commit('SET_SELECTED_ALERT_GROUP', null);
    },
  },
  modules: {}
});
