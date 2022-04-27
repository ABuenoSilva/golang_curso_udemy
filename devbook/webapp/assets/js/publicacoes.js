//Implementar descurtir com atualização de tela (um curtir por usuário), o do curso é um mock utilizando variáveis que se perdem qdo atualiza o path
$('#nova-publicacao').on('submit', criarPublicacao);
$(document).on('click', '.curtir-publicacao', curtirPublicacao); //usando document pq vai ter atualizar do html (não nessa versão, mas na do curso que atualizava a variável)
$('#atualizar-publicacao').on('click', atualizarPublicacao);
$('.deletar-publicacao').on('click', deletarPublicacao);

function criarPublicacao(event) {
    event.preventDefault();

    $.ajax({
        url: '/publicacoes',
        method: 'POST',
        data: {
            titulo: $('#titulo').val(),
            conteudo: $('#conteudo').val(),
        }
    }).done(function () {
        Swal.fire('Sucesso!', 'Publicação criada com sucesso', 'success')
            .then(() => window.location = '/home');        
    }).fail(function () {
        Swal.fire('Erro', 'Erro ao criar a publicação', 'error');
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
        Swal.fire('Erro', 'Erro ao curtir a publicação', 'error');
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
        Swal.fire('Sucesso!', 'Publicação atualizada com sucesso', 'success')
            .then(() => window.location = '/home');
    }).fail(function () {
        Swal.fire('Erro', 'Erro ao atualizar a publicação', 'error');
    }).always(function () {
        $('#atualizar-publicacao').prop('disabled', false);
    });
}

function deletarPublicacao(event) { 
    
    event.preventDefault();
    Swal.fire({
        title: "Atenção!",
        text: "Tem certeza que deseja exlcuir essa publicação? Essa ação é irreversível!",
        showCancelButton: true,
        cancelButtonText: "Cancelar",
        icon: "warning",
    }).then( (confirmar) => {
        if(confirmar.isConfirmed) {
            const elementoClicado = $(event.target);
            const publicacaoId = elementoClicado.closest('div').data('publicacao-id');
        
            $.ajax({
                url: `/publicacoes/${publicacaoId}`,
                method: "DELETE"
            }).done(function () {
                Swal.fire('Sucesso!', 'Publicação excluída com sucesso', 'success')
                    .then(() => window.location = '/home');
            }).fail(function () {
                Swal.fire('Erro', 'Erro ao excluir a publicação', 'error');
            }).always(function () {
                elementoClicado.prop('disabled', false);
            });                   
        };
    });

}