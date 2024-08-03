<script lang="ts">
</script>

<div class="flex gap-2 items-center">
    <span
        class="flex bg-black text-white rounded-full px-3 py-1.5 items-center justify-center"
        >GET</span
    >
    <p class="underline">https://uaas.tencorvids.com/right-pad</p>
</div>

<div class="flex flex-col gap-8">
    <p>
        <strong>Parameters:</strong><br />
        [<em>required</em>] <strong>string:</strong> The string to be padded<br
        />
        [<em>optional</em>] <strong>padding:</strong> The desired length after
        padding (default: 4)<br />
        [<em>optional</em>] <strong>padchar:</strong> The character to use for padding
        (default: space)
    </p>

    <p>
        <strong>Example Request:</strong><br />
        https://uaas.tencorvids.com/right-pad?string=hello&amp;padding=5&amp;padchar=-
    </p>

    <p>
        <strong>Example Response:</strong><br />
        hello-----
    </p>
</div>
