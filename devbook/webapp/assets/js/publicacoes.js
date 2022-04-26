//Implementar descurtir com atualização de tela (um curtir por usuário), o do curso é um mock utilizando variáveis que se perdem qdo atualiza o path
$('#nova-publicacao').on('submit', criarPublicacao);
$(document).on('click', '.curtir-publicacao', curtirPublicacao); //usando document pq vai ter atualizar do html (não nessa versão, mas na do curso que atualizava a variável)
$('#atualizar-publicacao').on('click', atualizarPublicacao);
$('.deletar-publicacao').on('click', deletarPublicacao);

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
        window.location = '/home';
    }).fail(function () {
        alert('Erro ao curtir publicação');
    }).always(function () {
        elementoClicado.prop('disabled', false);
    });

}

function atualizarPublicacao() { 
    
    $(this).prop('disabled', true);
    const publicacaoId = $(this).data('publicacao-id');
    $.ajax({
        url: `/publicacoes/${publicacaoId}`,
        method: "PUT",
        data: {
            titulo: $('#titulo').val(),
            conteudo: $('#conteudo').val()
        }
    }).done(function () {
        alert('Publicação atualizada com sucesso');
    }).fail(function () {
        alert('Erro ao atualizar publicação');
    }).always(function () {
        $('#atualizar-publicacao').prop('disabled', false);
    });
}

function deletarPublicacao(event) { 
    
    event.preventDefault();
    const elementoClicado = $(event.target);
    const publicacaoId = elementoClicado.closest('div').data('publicacao-id');

    $.ajax({
        url: `/publicacoes/${publicacaoId}`,
        method: "DELETE"
    }).done(function () {
        alert('Publicação excluída com sucesso!');
        window.location = '/home';
    }).fail(function () {
        alert('Erro ao excluir publicação');
    }).always(function () {
        elementoClicado.prop('disabled', false);
    });
}