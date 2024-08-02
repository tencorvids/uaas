<script lang="ts">
    import LeftPadDoc from "./LeftPadDoc.svelte";
    import RightPadDoc from "./RightPadDoc.svelte";
    import IsEvenDoc from "./IsEvenDoc.svelte";
    import IsOddDoc from "./IsOddDoc.svelte";
    import { onMount } from "svelte";

    interface Endpoint {
        name: string;
        id: string;
        component:
            | typeof LeftPadDoc
            | typeof RightPadDoc
            | typeof IsEvenDoc
            | typeof IsOddDoc;
    }

    let endpoints: Endpoint[] = [
        { name: "LEFT-PAD", id: "000_000_000", component: LeftPadDoc },
        { name: "RIGHT-PAD", id: "000_000_000", component: RightPadDoc },
        { name: "IS-EVEN", id: "000_000_000", component: IsEvenDoc },
        { name: "IS-ODD", id: "000_000_000", component: IsOddDoc },
    ];

    let selectedEndpoint: Endpoint | null = null;

    function selectEndpoint(endpoint: Endpoint) {
        if (selectedEndpoint === endpoint) {
            selectedEndpoint = null;
        } else {
            selectedEndpoint = endpoint;
        }
    }

    function handleKeydown(event: KeyboardEvent, endpoint: Endpoint) {
        if (event.key === "Enter" || event.key === " ") {
            event.preventDefault();
            selectEndpoint(endpoint);
        }
    }

    function formatCount(count: number): string {
        return count
            .toString()
            .padStart(9, "0")
            .replace(/(\d{3})(?=\d)/g, "$1_");
    }

    onMount(async () => {
        try {
            const response = await fetch("http://localhost:8080/counts");
            const counts = await response.json();
            endpoints = endpoints.map((endpoint) => ({
                ...endpoint,
                id: formatCount(counts[endpoint.name.toLowerCase()] || 0),
            }));
        } catch (error) {
            console.error("Failed to fetch endpoint counts:", error);
        }
    });
</script>

<div
    class="w-full h-full flex flex-col justify-between border-t-2 border-t-black pt-4"
>
    <div class="flex flex-col gap-16">
        <ul class="flex flex-col gap-2 list-none p-0" role="list">
            {#each endpoints as endpoint}
                <li>
                    <button
                        class="flex w-full justify-between cursor-pointer hover:bg-gray-100 p-1 text-left"
                        on:click={() => selectEndpoint(endpoint)}
                        on:keydown={(event) => handleKeydown(event, endpoint)}
                        aria-pressed={selectedEndpoint === endpoint}
                    >
                        <span
                            class={selectedEndpoint === endpoint
                                ? "font-medium"
                                : ""}
                        >
                            {endpoint.name}
                        </span>
                        <span>{endpoint.id}</span>
                    </button>
                </li>
            {/each}
        </ul>

        {#if selectedEndpoint}
            <svelte:component this={selectedEndpoint.component} />
        {/if}
    </div>

    <div class="w-full flex justify-center">
        <slot name="barcode" />
    </div>
</div>

