$('#parar-de-seguir').on('click', pararDeSeguir);
$('#seguir').on('click', seguir);

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