// assets/js/script.js
let tarefasGlobal = [];
let idEdicao = null;

async function carregarTarefas() {
    try {
        const resposta = await fetch('/tarefas');
        const tarefas = await resposta.json();
        tarefasGlobal = tarefas;

        const tabela = document.getElementById('tabela-corpo');
        tabela.innerHTML = '';
        
        let total = 0;

        tarefas.forEach(t => {
            total += parseFloat(t.custo); 

            const custoFormatado = parseFloat(t.custo).toLocaleString('pt-BR', { style: 'currency', currency: 'BRL' });
            //fuso horario tava retornando um dia antes
            const dataObj = new Date(t.data_limite);
            dataObj.setMinutes(dataObj.getMinutes() + dataObj.getTimezoneOffset());
            const dataFormatada = dataObj.toLocaleDateString('pt-BR');
            
            const classeLinha = t.custo >= 1000 ? 'table-warning' : '';

            const linha = `
                <tr class="${classeLinha}">
                    <td>${t.ordem_apresentacao}</td>
                    <td>${t.nome}</td>
                    <td>${custoFormatado}</td>
                    <td>${dataFormatada}</td>
                    <td>
                        <button class="btn btn-sm btn-warning" onclick="editarTarefa(${t.id})">✏️</button>
                        <button class="btn btn-sm btn-danger" onclick="excluirTarefa(${t.id})">🗑️</button>
                        <button class="btn btn-sm btn-secondary" onclick="moverTarefa(${t.id}, 'subir')">⬆️</button>
                        <button class="btn btn-sm btn-secondary" onclick="moverTarefa(${t.id}, 'descer')">⬇️</button>
                    </td>
                </tr>
            `;
            tabela.innerHTML += linha;
        });

        document.getElementById('custo-total').innerText = total.toLocaleString('pt-BR', { minimumFractionDigits: 2 });

    } catch (erro) {
        console.error("Erro ao carregar:", erro);
    }
}

function abrirModalIncluir() {
    idEdicao = null;
    document.getElementById('tituloModal').innerText = "Nova Tarefa";
    document.getElementById('nome').value = '';
    document.getElementById('custo').value = '';
    document.getElementById('dataLimite').value = '';
    
    const modal = new bootstrap.Modal(document.getElementById('modalTarefa'));
    modal.show();
}

async function salvarTarefa() {
    const nome = document.getElementById('nome').value;
    const custo = document.getElementById('custo').value;
    const dataLimite = document.getElementById('dataLimite').value;

    if (!nome || !custo || !dataLimite) {
        mostrarNotificacao("Por favor, preencha todos os campos antes de salvar!", "erro");
        return; 
    }

    const novaTarefa = {
        nome: nome,
        custo: parseFloat(custo),
        data_limite: new Date(dataLimite).toISOString()
    };

    const metodo = idEdicao ? 'PUT' : 'POST';
    const url = idEdicao ? `/tarefas/${idEdicao}` : '/tarefas';

    try {
        const resposta = await fetch(url, {
            method: metodo,
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify(novaTarefa)
        });

        if (resposta.ok) {
            mostrarNotificacao(idEdicao ? "Tarefa editada com sucesso!" : "Tarefa incluída com sucesso!", "sucesso");
            setTimeout(() => location.reload(), 1500); //o reload é muito rapido, o toast não tava aparecendo então botei um delay
        } else {
            const erro = await resposta.json();
            mostrarNotificacao("Erro: " + erro.erro, "erro");
        }
    } catch (e) {
        console.error(e);
        mostrarNotificacao("Erro ao salvar a tarefa!", "erro");
    }
}

function editarTarefa(id) {
    
    const tarefa = tarefasGlobal.find(t => t.id === id);
    if(tarefa) {
        idEdicao = id;
        document.getElementById('tituloModal').innerText = "Editar Tarefa";
        document.getElementById('nome').value = tarefa.nome;
        document.getElementById('custo').value = tarefa.custo;
        document.getElementById('dataLimite').value = tarefa.data_limite.split('T')[0];
        
        const modal = new bootstrap.Modal(document.getElementById('modalTarefa'));
        modal.show();
    }
}


async function excluirTarefa(id) {
    const confirmacao = confirm("Tem certeza que deseja excluir?");
    if (!confirmacao) return;

    try {
        const resposta = await fetch(`/tarefas/${id}`, { method: 'DELETE' });

        if (resposta.ok) {
            mostrarNotificacao("Tarefa excluída!", "sucesso");
            setTimeout(() => location.reload(), 1500);
        } else {
            const erro = await resposta.json();
            mostrarNotificacao("Erro: " + erro.erro, "erro");
        }
    } catch (e) {
        console.error(e);
        mostrarNotificacao("Erro ao excluir!", "erro");
    }
}

async function moverTarefa(id, direcao) {
    try {
        const resposta = await fetch(`/tarefas/${id}/${direcao}`, { method: 'PUT' });
        if (resposta.ok) {
            mostrarNotificacao("Ordem alterada com sucesso!", "sucesso");
            setTimeout(() => location.reload(), 1500);
        } else {
            const erro = await resposta.json();
            mostrarNotificacao("Erro ao mover: " + (erro.erro || "Desconhecido"), "erro");
        }
    } catch (e) {
        console.error(e);
        mostrarNotificacao("Erro de conexão!", "erro");
    }
}

function mostrarNotificacao(mensagem, tipo = 'sucesso') {
    const toastEl = document.getElementById('liveToast');
    const corpoToast = document.getElementById('toast-mensagem');
    corpoToast.innerText = mensagem;
    toastEl.className = `toast align-items-center text-white border-0 ${tipo === 'erro' ? 'bg-danger' : 'bg-success'}`;
    const toast = new bootstrap.Toast(toastEl);
    toast.show();
}

carregarTarefas();