<script>
	import { onDestroy, onMount } from "svelte";
	import FolderView from "./Components/FolderView.svelte";
	import Sidebar from "./Components/Sidebar.svelte";

	const basePath = window.BASE_PATH || "/";
	let path = location.pathname.replace(basePath, '');

	function navigate(newPath) {
		path = newPath;
	}

	function onNavigate(e) {
		navigate(e.detail.path);
		history.pushState(
			{ path },
			"",
			encodeURI(basePath + (path === "" ? "" : path + "/"))
		);
	}

	function onStateChange(e) {
		const { state } = e;
		if (state && state.path) {
			navigate(state.path);
		}
	}

	onMount(() => window.addEventListener("popstate", onStateChange));
	onDestroy(() => window.removeEventListener("popstate", onStateChange));
</script>

<div class="root">
	<header class="header">
		<div class="title">Files</div>
	</header>
	<div class="wrapper">
		<Sidebar />
		<div class="main">
			<FolderView {path} on:navigate={onNavigate} />
		</div>
	</div>
</div>

<style>
	:global(html, body) {
		margin: 0;
		padding: 0;
		width: 100%;
		height: 100%;
	}

	.root {
		display: flex;
		flex-direction: column;
		width: 100%;
		height: 100%;
		overflow: hidden;
	}

	.header {
		height: 48px;
		line-height: 48px;
		padding: 0 16px;
		color: #ffffff;
		background-color: #3376cd;
		user-select: none;
	}

	.header .title {
		font-size: 20px;
		font-weight: normal;
	}

	.wrapper {
		flex: 1;
		display: flex;
		flex-direction: row;
		overflow: hidden;
	}

	.wrapper .main {
		flex: 1;
		overflow: auto;
		padding: 16px;
	}
</style>
