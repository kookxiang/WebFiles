<script>
    import bytes from "bytes";
    import moment from "moment";
    import { createEventDispatcher } from "svelte";
    import Breadcrumb from "./Breadcrumb.svelte";

    export let path = "";
    let files = [];

    const dispatch = createEventDispatcher();
    const basePath = window.BASE_PATH || "/";

    $: loadFiles(path);

    async function loadFiles(relativePath) {
        files = [];
        const response = await fetch(basePath + relativePath, {
            headers: { Accept: "application/json" },
        });
        const data = await response.json();
        data.sort((a, b) => {
            if (b.directory !== a.directory) {
                return b.directory - a.directory;
            }
            return a.name.localeCompare(b.name);
        });
        if (path === relativePath) {
            files = data;
        }
    }

    function onPreviewFile(file) {
        const target = (path ? path + "/" : "") + file.name;
        if (file.directory) {
            dispatch("navigate", { path: target });
            return;
        } else {
            const el = document.createElement("a");
            el.href = file.name;
            el.download = file.name;
            el.click();
        }
    }
</script>

<div class="folder">
    <Breadcrumb {path} on:navigate />
    <!-- <Toolbar /> -->
    <div class="file-list">
        <div class="file-header">
            <div class="name">文件名</div>
            <div class="size">大小</div>
            <div class="modify-time">修改时间</div>
        </div>
        {#each files as file}
            <div
                on:click={() => onPreviewFile(file)}
                on:contextmenu|preventDefault={() => console.log(file)}
                class="file-item"
            >
                {#if file.directory}
                    <div class="file-icon folder" />
                {:else}
                    <div class="file-icon ext-{file.name.split('.').pop()}" />
                {/if}
                <div class="name">{file.name}</div>
                <div class="size">
                    {#if !file.directory}
                        {bytes(file.size, { unitSeparator: " " })}
                    {/if}
                </div>
                <div class="modify-time">
                    <span
                        title={moment(file.mtime).format("YYYY-MM-DD HH:mm:ss")}
                    >
                        {moment(file.mtime).fromNow()}
                    </span>
                </div>
            </div>
        {/each}
    </div>
</div>

<style>
    .file-list {
        min-width: 100%;
        font-size: 14px;
    }

    .file-header {
        display: flex;
        padding-left: 36px;
        margin: 16px auto 8px;
    }

    .file-header > * {
        padding: 10px;
        cursor: pointer;
        transition: all .2s ease-out;
    }

    .file-header > *:hover {
        background-color: rgba(0, 0, 0, 0.1);
    }

    .file-header .name {
        flex: 1;
    }

    .file-header .size {
        width: 60px;
        text-align: center;
    }

    .file-header .modify-time {
        width: 100px;
        text-align: center;
    }

    .file-item {
        display: flex;
        border: 1px solid transparent;
        margin: -2px 0;
        padding: 8px 0 8px 10px;
        line-height: 24px;
        cursor: pointer;
        transition: all 0.2s ease-in;
    }

    .file-item:hover {
        background-color: #f5f5f5;
        border: 1px solid #cccccc;
    }

    .file-item > .file-icon {
        width: 24px;
        height: 24px;
        margin-top: -1px;
        margin-right: 10px;
    }

    .file-item > .name {
        flex: 1;
        overflow: hidden;
        white-space: nowrap;
        text-overflow: ellipsis;
    }

    .file-item > .size {
        width: 80px;
        text-align: center;
    }

    .file-item > .modify-time {
        width: 120px;
        text-align: center;
    }

    .file-icon {
        background-image: url(../Resources/Icon/file-empty.svg);
        background-position: 50% 50%;
        background-repeat: no-repeat;
        background-size: 24px;
    }

    .file-icon.folder {
        background-image: url(../Resources/Icon/folder.svg);
    }

    .file-icon.ext-js,
    .file-icon.ext-json,
    .file-icon.ext-css,
    .file-icon.ext-less,
    .file-icon.ext-scss {
        background-image: url(../Resources/Icon/file-code.svg);
    }

    .file-icon.ext-xls,
    .file-icon.ext-xlsx {
        background-image: url(../Resources/Icon/file-excel.svg);
    }

    .file-icon.ext-doc,
    .file-icon.ext-docx {
        background-image: url(../Resources/Icon/file-word.svg);
    }

    .file-icon.ext-ppt,
    .file-icon.ext-pptx {
        background-image: url(../Resources/Icon/file-powerpoint.svg);
    }

    .file-icon.ext-7z,
    .file-icon.ext-tar,
    .file-icon.ext-gz,
    .file-icon.ext-rar,
    .file-icon.ext-zip {
        background-image: url(../Resources/Icon/file-zip.svg);
    }

    .file-icon.ext-pdf {
        background-image: url(../Resources/Icon/file-pdf.svg);
    }

    .file-icon.ext-dll,
    .file-icon.ext-exe {
        background-image: url(../Resources/Icon/file-exe.svg);
    }

    .file-icon.ext-png,
    .file-icon.ext-gif,
    .file-icon.ext-tif,
    .file-icon.ext-tiff,
    .file-icon.ext-jpg,
    .file-icon.ext-jpeg {
        background-image: url(../Resources/Icon/file-picture.svg);
    }

    .file-icon.ext-avi,
    .file-icon.ext-mov,
    .file-icon.ext-mkv,
    .file-icon.ext-mp4 {
        background-image: url(../Resources/Icon/file-video.svg);
    }

    .file-icon.ext-aac,
    .file-icon.ext-flac,
    .file-icon.ext-wav,
    .file-icon.ext-m4a,
    .file-icon.ext-mp3 {
        background-image: url(../Resources/Icon/file-sound.svg);
    }
</style>
