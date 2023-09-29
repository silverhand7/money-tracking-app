import { createRouter, createWebHistory } from 'vue-router'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/login',
      name: 'login',
      component: () => import('@/views/LoginView.vue'),
    },
    {
      path: '/',
      name: 'home',
      alias: ['/index.html'],
      // component: HomeView
      component: () => import('@/views/CategoryView.vue'),
      meta: {
        requiresAuth: true
      }
    },
    {
      path: '/about',
      name: 'about',
      // route level code-splitting
      // this generates a separate chunk (About.[hash].js) for this route
      // which is lazy-loaded when the route is visited.
      component: () => import('../views/AboutView.vue')
    },
    {
      path: '/categories',
      name: 'categories',
      component: () => import('@/views/CategoryView.vue'),
      meta: {
        requiresAuth: true
      }
    },
    {
      path: '/categories/create',
      name: 'categories.create',
      component: () => import('@/views/CategoryForm.vue'),
      meta: {
        requiresAuth: true
      }
    },
    {
      path: '/wallets',
      name: 'wallets',
      component: () => import('@/views/WalletView.vue'),
      meta: {
        requiresAuth: true
      }
    },
    {
      path: '/wallet/create',
      name: 'wallets.create',
      component: () => import('@/views/WalletForm.vue'),
      meta: {
        requiresAuth: true
      }
    },
    {
      path: '/user',
      name: 'user',
      component: () => import('@/views/UserView.vue'),
      meta: {
        requiresAuth: true
      }
    },
    {
      path: '/:pathMatch(.*)*',
      name: 'NotFound',
      component: () => import('@/views/NotFoundView.vue'),
      meta: {
        requiresAuth: true
      }
    }
  ]
})

router.beforeEach((to, from, next) => {
  if (to.meta.requiresAuth) {
    const user = localStorage.getItem('user');
    if (user) {
      // User is authenticated, proceed to the route
      next();
    } else {
      // User is not authenticated, redirect to login
      next('/login');
    }
  } else {
    next();
  }
});

export default router
