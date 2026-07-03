import { createRouter, createWebHistory } from 'vue-router'
import DashboardView from '../views/DashboardView.vue'
import { useAuthStore } from '@/stores/auth'

const router = createRouter({
    history: createWebHistory(import.meta.env.BASE_URL),
    routes: [
        {
            path: '/',
            name: 'dashboard',
            component: DashboardView,
            meta: { requiresAuth: true }
        },
        {
            path: '/login',
            name: 'login',
            component: () => import('../views/LoginView.vue'),
            meta: { requiresGuest: true },
        },
        {
            path: '/users',
            name: 'users',
            component: () => import('../views/UsersView.vue'),
            meta: { requiresAuth: true, requiresRole: ['superadmin'] },
            children: [
                {
                    path: '',
                    name: 'users.list',
                    component: () => import('../components/users/Users.vue'),
                },
                {
                    path: 'create',
                    name: 'users.create',
                    component: () => import('../components/users/UsersCreate.vue'),
                },
                {
                    path: ':id/edit',
                    name: 'users.edit',
                    component: () => import('../components/users/UsersEdit.vue'),
                },
                {
                    path: ':id',
                    name: 'users.show',
                    component: () => import('../components/users/UsersShow.vue'),
                },
            ],
        },
        {
            path: '/activities',
            name: 'activities',
            component: () => import('../views/ActivitiesView.vue'),
            meta: { requiresAuth: true },
            children: [
                {
                    path: '',
                    name: 'activities.list',
                    component: () => import('../components/activities/Activities.vue'),
                },
                {
                    path: 'create',
                    name: 'activities.create',
                    component: () => import('../components/activities/ActivitiesCreate.vue'),
                    meta: { requiresRole: ['admin', 'superadmin'] }
                },
                {
                    path: ':id/edit',
                    name: 'activities.edit',
                    component: () => import('../components/activities/ActivitiesEdit.vue'),
                    meta: { requiresRole: ['admin', 'superadmin'] }
                },
            ]
        },
        {
            path: '/reports',
            name: 'reports',
            component: () => import('../views/ReportsView.vue'),
            meta: { requiresAuth: true, requiresRole: ['admin', 'superadmin'] },
            children: [
                {
                    path: 'kegiatan',
                    name: 'reports.kegiatan',
                    component: () => import('../components/reports/ReportsKegiatan.vue'),
                },
            ]
        },
        {
            path: '/organization',
            name: 'organization',
            component: () => import('../views/OrganizationView.vue'),
            meta: { requiresAuth: true, requiresRole: ['superadmin'] },
            children: [
                {
                    path: '',
                    name: 'organization.list',
                    component: () => import('../components/organization/Organization.vue'),
                },
                {
                    path: 'create',
                    name: 'organization.create',
                    component: () => import('../components/organization/OrganizationCreate.vue'),
                },
                {
                    path: 'edit/:id',
                    name: 'organization.edit',
                    component: () => import('../components/organization/OrganizationEdit.vue'),
                }
            ]
        },
        {
            path: '/program',
            name: 'program',
            component: () => import('../views/ProgramView.vue'),
            meta: { requiresAuth: true, requiresRole: ['superadmin'] },
            children: [
                {
                    path: '',
                    name: 'program.list',
                    component: () => import('../components/program/Program.vue'),
                },
                {
                    path: 'create',
                    name: 'program.create',
                    component: () => import('../components/program/ProgramCreate.vue'),
                },
                {
                    path: 'edit/:id',
                    name: 'program.edit',
                    component: () => import('../components/program/ProgramEdit.vue'),
                }
            ]
        },
    ],
});

import { isTokenValid } from '@/utils/cek_token'

router.beforeEach((to, from, next) => {
    const isAuthenticated = isTokenValid();
    const auth = useAuthStore();
    const role = auth.user?.role;

    if (to.meta.requiresAuth && !isAuthenticated) {
        next({ name: 'login' });
    } else if (to.meta.requiresGuest && isAuthenticated) {
        next({ name: 'dashboard' });
    } else if (to.name === 'users.show' && to.params.id == auth.user?.id) {
        next();
    } else if (to.meta.requiresRole && !to.meta.requiresRole.includes(role)) {
        next({ name: 'dashboard' });
    } else {
        next();
    }
});

export default router
