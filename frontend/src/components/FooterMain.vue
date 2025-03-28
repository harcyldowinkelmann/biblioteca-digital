<template>
    <div class="versaoSistema">
        <span class="text-footer"> Biblioteca Digital 2025 - Versão: <a class="link">{{ pack.version }} </a> | Projeto de Sistemas Computacionais II - UFGD | Copyright 2024 Todos os direitos reservados &copy;</span>
    </div>
</template>

<script>

export default {
    name: "FooterView",
    data: () => ({
        pack: require('../../package.json')
    })
}
</script>


<style scoped>

.versaoSistema {
    text-align: center;
    margin-top: 2rem;
    font-weight: bold;
    font-size: small;
}

.link {
    text-decoration: underline;
    font-weight: bold;
}

.text-footer {
    color: white;
    font-weight: bolder;
}

</style>
