import { createRouter, createWebHistory } from 'vue-router';
import { queryClient } from '@/query';
import LoginView from '@/app/views/LoginView.vue';
import RegisterView from '@/app/views/RegisterView.vue';

const routing = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'dashboard',
      component: () => import('@/app/views/DashboardView.vue'),
      meta: { requiresAuth: true },
    },
    {
      path: '/login',
      name: 'login',
      component: LoginView,
    },
    {
      path: '/sign-up',
      name: 'sign-up',
      component: RegisterView,
    },
  ],
});

routing.beforeEach(async (to) => {
  const user = queryClient.getQueryData(['auth']); // 🔁 swap with your actual query key

  const isAuthenticated = !!user;

  if (isAuthenticated && to.name === 'login') {
    return { name: 'dashboard' };
  }

  if (!isAuthenticated && to.meta.requiresAuth) {
    return { name: 'login' };
  }
});

export default routing;
