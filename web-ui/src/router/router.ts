import { createRouter, createWebHistory, RouteRecordRaw } from "vue-router";
import DashboardView from "../views/DashboardView.vue";
import DeviceConfigView from "../views/DeviceConfigView.vue";

const routes: Array<RouteRecordRaw> = [
  {
    path: "/",
    redirect: "/dashboard",
  },

  {
    path: "/dashboard",
    name: "dashboard",
    component: DashboardView,
  },
  {
    path: "/deviceConfig",
    name: "deviceConfig",
    component: DeviceConfigView,
  },
];

const router = createRouter({
  history: createWebHistory(),
  routes: routes,
});

export default router;
