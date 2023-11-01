<script>
    import { metatags } from '@roxi/routify'
    import { g_get_crypto_session, g_set_crypto_session, set_session_crypto_key } from "../js/common/mattabu_common.js";
    import { onMount } from 'svelte'
    
    metatags.title = 'My Routify app'
    metatags.description = 'Description coming soon...'

    onMount(() => {
        set_session_crypto_key("ABC");
        g_set_crypto_session("USER_ID", "mattabu");
        const user_id = g_get_crypto_session("USER_ID");
        console.log("USER_ID:", user_id);
    });
</script>

<h1>Routify Starter Template</h1>

<p>
    To see an example app that shows off a lot of Routify's features, go to <a href="/example">/example</a>
</p>

<p>
    This template is ready to be used in production! just delete the example app at: src/pages/example
</p>