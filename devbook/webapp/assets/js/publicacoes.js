//Implementar descurtir com atualização de tela (um curtir por usuário), o do curso é um mock utilizando variáveis que se perdem qdo atualiza o path
$('#nova-publicacao').on('submit', criarPublicacao);
$('.curtir-publicacao').on('click', curtirPublicacao);

function criarPublicacao(event) {
    event.preventDefault;

    $.ajax({
        url: '/publicacoes',
        method: 'POST',
        data: {
            titulo: $('#titulo').val(),
            conteudo: $('#conteudo').val(),
        }
    }).done(function() {
        window.location='/home';
    }).fail(function () {
        alert('Erro ao criar a publicação');
    })
}

function curtirPublicacao(event) {    
    event.preventDefault();
    const elementoClicado = $(event.target);
    const publicacaoId = elementoClicado.closest('div').data('publicacao-id');
    elementoClicado.prop('disabled', true);
    $.ajax({
        url: `/publicacoes/${publicacaoId}/curtir`,
        method: "POST"
    }).done(function () {
        window.location = '/home';;
    }).fail(function () {
        alert('Erro ao curtir publicação');
    }).always(function () {
        elementoClicado.prop('disabled', false);
    });

}