<script>
    import BreadcrumbItem from "./BreadcrumbItem.svelte";

    export let path = "";
    let parts;

    $: parts = path.split("/").filter(Boolean);
</script>

<div class="breadcrumb">
    <BreadcrumbItem home {parts} on:navigate />
    {#each parts as _, i}
        <BreadcrumbItem index={i} {parts} on:navigate />
    {/each}
</div>

<style>
    .breadcrumb {
        display: flex;
        height: 48px;
        align-items: center;
    }
</style>
