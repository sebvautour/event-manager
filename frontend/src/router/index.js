import Vue from "vue";
import VueRouter from "vue-router";
import AlertConsole from "../views/AlertConsole.vue";

Vue.use(VueRouter);

const routes = [
  {
    path: "/",
    name: "Console",
    component: AlertConsole
  }
];

const router = new VueRouter({
  mode: "history",
  base: process.env.BASE_URL,
  routes
});

export default router;
