import {LitElement, html, render} from 'https://cdn.jsdelivr.net/gh/lit/dist@3/core/lit-core.min.js';

const app = () => html`
    <div class="container">
        <h1>Hello!! </h1>
        <h1>Golang + Postgres + Lit Element + Docker</h1>
    </div>
`;
render(app(), document.getElementById("root"));