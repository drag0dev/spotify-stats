<style>
    *{
        margin: 0;
    }
    .index{
        display: flex;
        flex-wrap: wrap;
        flex-direction: row;
        margin-left: 20vw;
        margin-right: 20vw;
        margin-top: 0;
        margin-bottom: 0;
        width: 60vw;
    }

    .options{
        width: 100%;
    }

    .options{
            display: flex;
            flex-wrap: wrap;
            flex-direction: row;
            justify-items: center;
            justify-content: center;
            align-items: center;
            height: 20vh;
        }
    .options button{
        width: 20%;
        height: 25%;

        color: #1DB954;
        background-color: black;
        border: 2px solid #1DB954;
        box-shadow: 3px 3px 3px 1px #1DB954;
        border-radius: 20px;
        cursor: pointer;
    }
    .options button:hover{
        border: 2px solid lightblue;
        box-shadow: 3px 3px 3px 1px lightblue;
        color: lightblue;
    }
</style>

<script>
    import Footer from "./_components/Footer.svelte";
    import Header from "./_components/Header.svelte";
    import auth from "../lib/auth";
</script>

<div class="index">
    <Header />
    
    <div class="options" on:click={auth}>
            <button>Log in</button>
    </div>

    <Footer />
</div>
