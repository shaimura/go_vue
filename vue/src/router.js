import Vue from 'vue';
import Router from 'vue-router';

const Home = () => import( /* webpackChunkName: "Home" */ './views/Home.vue');
const Top = () => import( /* webpackChunkName: "Top" */ './components/Top.vue');
const Header = () => import( /* webpackChunkName: "Header" */ './components/Header.vue');
const Footer = () => import( /* webpackChunkName: "Footer" */ './components/Footer.vue');
const About = () => import( /* webpackChunkName: "About" */ './components/About.vue');
const Skill = () => import( /* webpackChunkName: "Skill" */ './components/Skill.vue');
const Portfolio = () => import( /* webpackChunkName: "Portfolio" */ './components/Portfolio.vue');
const Contact = () => import( /* webpackChunkName: "Contact" */ './components/Contact.vue');
const Login = () => import( /* webpackChunkName: "Login" */ './views/Login.vue');
const Sinup = () => import( /* webpackChunkName: "Sinup" */ './views/Sinup.vue');
const Users = () => import( /* webpackChunkName: "Users" */ './views/Users.vue');
const Userchat = () => import( /* webpackChunkName: "Userchat" */ './views/Userchat.vue');


Vue.use(Router);

export default new Router({
    mode: 'history',
    routes: [
        {
            path: '/',
            components: {
                default: Home,
                Top: Top,
                About: About,
                Skill: Skill,
                Portfolio: Portfolio,
                Contact: Contact,
                Header: Header,
                Footer: Footer
                // FadeIn: FadeIn
            }
        },
        {
            path: '/login',
            component: Login
        },
        {
            path: '/sinup',
            component: Sinup
        },
        {
            path: '/users',
            component: Users
        },
        {
            path: '/:firstuser/:firstid/:seconduser/:secondid/chatroom',
            component: Userchat,
            props: true
        },
        {
        path: '*',
        redirect: '/'
        }
    ]
    
});