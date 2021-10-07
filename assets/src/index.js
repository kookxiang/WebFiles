import 'normalize.css';
import App from './App.svelte';
import moment from 'moment';
import 'moment/locale/zh-cn';

moment.locale('zh-cn');

var app = new App({ target: document.body });

export default app;