$('#parar-de-seguir').on('click', pararDeSeguir);
$('#seguir').on('click', seguir);
$('#editar-usuario').on('submit', editar);
$('#atualizar-senha').on('submit', atualizarSenha);
$('#deletar-usuario').on('click', deletarUsuario);

function pararDeSeguir(event) {
    event.preventDefault();
    const usuarioId = $(this).data('usuario-id');

    $(this).prop('disabled', true);
    
    $.ajax({
        url: `/usuarios/${usuarioId}/parar-de-seguir`,
        method: 'POST'
    }).done(function () {
        window.location.reload();
    }).fail(function () {
        Swal.fire('Ops...', 'Falha ao parar de seguir o usuário!', 'error');
        $('#parar-de-seguir').prop('disabled', false);
    });
}

function seguir(event) {
    event.preventDefault();
    const usuarioId = $(this).data('usuario-id');

    $(this).prop('disabled', true);
    
    $.ajax({
        url: `/usuarios/${usuarioId}/seguir`,
        method: 'POST'
    }).done(function () {
        window.location.reload();
    }).fail(function () {
        Swal.fire('Ops...', 'Falha ao seguir o usuário!', 'error');
        $('#parar-de-seguir').prop('disabled', false);
    });
}

function editar(event) {
    event.preventDefault();

    $.ajax({
        url: '/editar-usuario',
        method: 'PUT',
        data: {
            nome: $('#nome').val(),
            email: $('#email').val(),
            nick: $('#nick').val()
        }
    }).done(function () {
        Swal.fire('Sucesso!', 'Usuário atualizado com sucesso', 'success')
            .then(() => window.location = '/perfil');
    }).fail(function () {
        Swal.fire('Ops!', 'Erro ao atualizar o usuário', 'error')
    })
}

function atualizarSenha(event) {
    event.preventDefault();
    if ($('#nova-senha').val() != $('#confirmar-senha').val()) {
        Swal.fire('Ops!', 'As senhas não coincidem', 'warning');
        return;
    }

    $.ajax({
        url: '/atualizar-senha',
        method: 'POST',
        data: {
            atual: $('#senha-atual').val(),
            nova: $('#nova-senha').val()
        }
    }).done(function () {
        Swal.fire('Sucesso!', 'Senha atualizada com sucesso', 'success')
            .then(() => window.location = '/perfil');
    }).fail(function () {
        Swal.fire('Ops!', 'Erro ao atualizar a senha', 'error')
    });

}

function deletarUsuario(event) {
    Swal.fire({
        title: 'Atenção!',
        text: 'Tem certeza que deseja apagar a sua conta? Essa é uma ação irreversível!',
        showCancelButton: true,
        cancelButtonText: 'Cancelar',
        icon: 'warning'
    }).then(function (confirmar) {
        if (confirmar.isConfirmed) {
            $.ajax({
                url: '/deletar-usuario',
                method: 'DELETE'
            }).done(function () {
                Swal.fire('Sucesso!', 'Seu usuário foi excluído com sucesso', 'success')
                    .then(window.location = '/logout');
            }).fail(function () {
                Swal.fire('Ops!', 'Erro ao excluir o usuário', 'error');
            });
        }
    })
}