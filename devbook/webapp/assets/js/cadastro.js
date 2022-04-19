$('#formulario-cadastro').on('submit', criarUsuario)

function criarUsuario(event) {
    event.preventDefault();

    if ($('#senha').val() !== $('#confirmar-senha').val()) {
        alert("As senhas não conferem");
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
        alert('Usuário cadastrado com sucesso');
    }).fail( function(erro) {
        alert('Erro ao cadastrar o usuário: ' + erro.responseJSON.erro);
    });
}