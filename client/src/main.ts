import './assets/main.css';
import { createApp } from 'vue';
import { VueQueryPlugin, QueryClient } from '@tanstack/vue-query';
import App from './App.vue';
import routing from './router/routing';
const queryClient = new QueryClient();
const app = createApp(App);

app.use(routing).use(VueQueryPlugin, { queryClient });

app.mount('#app');
