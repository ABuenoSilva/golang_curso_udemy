$('#formulario-cadastro').on('submit', criarUsuario)

function criarUsuario(event) {
    event.preventDefault();

    if ($('#senha').val() !== $('#confirmar-senha').val()) {
        Swal.fire('Erro','As senhas não conferem','error');
        return;
    } 

    $.ajax({
        url: "/usuarios",
        method: "POST",
        data: {
            nome: $('#nome').val(),
            email: $('#email').val(),
            nick: $('#nick').val(),
            senha: $('#senha').val()
        }
    }).done( function() {
        Swal.fire('Sucesso!', 'Usuário cadastrado com sucesso', 'success')
            .then(() => {
                $.ajax({
                    url: "/login",
                    method: "POST",
                    data: {
                        email: $('#email').val(),
                        senha: $('#senha').val()
                    }
                }).done(function () {
                    window.location = '/home';
                }).fail(function () {
                    Swal.fire('Ops','Erro ao autenticar o usuário.','error');
                });

            });
    }).fail(function (erro) {
        Swal.fire('Ops','Erro ao cadastrar o usuário: ' + erro.responseJSON.erro,'error');
    });
}