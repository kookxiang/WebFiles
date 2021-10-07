<script>
    import { createEventDispatcher } from "svelte";

    export let home = false;
    export let index = 0;
    export let parts = [];

    const basePath = window.BASE_PATH || '/';
    const dispatch = createEventDispatcher();
    let path;

    $: path = home ? "" : parts.slice(0, index + 1).join("/");

    function onClick() {
        dispatch("navigate", { path });
    }
</script>

<div class="item">
    <a href={basePath+path} on:click|preventDefault={onClick}>
        {#if !home}
            {parts[index]}
        {/if}
    </a>
</div>

<style>
    .item {
        display: inline-block;
    }

    .item a {
        display: inline-block;
        border-radius: 4px;
        padding: 0 6px;
        height: 28px;
        line-height: 28px;
        color: #666;
        white-space: nowrap;
        text-decoration: none;
        vertical-align: middle;
        transition: all 0.5s ease-out;
    }

    .item a:hover {
        background-color: rgba(0, 0, 0, 0.1);
    }

    .item:first-child a {
        margin-top: -2px;
        background: none;
    }

    .item:first-child a:after {
        background: url(../Resources/Home.svg) no-repeat center center;
        background-size: 24px;
        content: "";
        display: inline-block;
        width: 24px;
        height: 24px;
        opacity: 0.5;
        vertical-align: middle;
    }

    .item:last-child a {
        color: #000;
        font-weight: bold;
    }

    .item:before {
        content: ">";
        display: inline-block;
        line-height: 28px;
        margin: -2px 4px 0;
        color: #666;
        vertical-align: middle;
    }

    .item:first-child:before {
        display: none;
    }
</style>
