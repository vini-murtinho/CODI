// Interface para Task (exportada para reutilização)
export interface Task {
  id: string;
  title: string;
  description?: string;
  completed: boolean;
  status: "todo" | "in_progress" | "done";
}

// Tipos de status das colunas
export type ColumnType = "todo" | "in_progress" | "done";

// URL base da API do backend
const API_BASE_URL = "http://localhost:8080";

// Função para fazer requisição GET de todas as tarefas
export const fetchAllTasks = async (): Promise<Task[]> => {
  try {
    const response = await fetch(`${API_BASE_URL}/tasks`);
    if (!response.ok) {
      throw new Error(`Erro ao buscar tarefas: ${response.statusText}`);
    }
    return await response.json();
  } catch (error) {
    console.error("Erro ao buscar tarefas:", error);
    return [];
  }
};

// Função para criar uma nova tarefa
export const createTask = async (
  title: string,
  description?: string
): Promise<Task | null> => {
  try {
    const response = await fetch(`${API_BASE_URL}/tasks`, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        title,
        description: description || "",
      }),
    });
    if (!response.ok) {
      throw new Error(`Erro ao criar tarefa: ${response.statusText}`);
    }
    return await response.json();
  } catch (error) {
    console.error("Erro ao criar tarefa:", error);
    return null;
  }
};

// Função para atualizar uma tarefa
export const updateTask = async (
  id: string,
  updates: Partial<{
    title: string;
    description: string;
    status: ColumnType;
  }>
): Promise<Task | null> => {
  try {
    const response = await fetch(`${API_BASE_URL}/tasks/${id}`, {
      method: "PUT",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(updates),
    });
    if (!response.ok) {
      throw new Error(`Erro ao atualizar tarefa: ${response.statusText}`);
    }
    return await response.json();
  } catch (error) {
    console.error("Erro ao atualizar tarefa:", error);
    return null;
  }
};

// Função para deletar uma tarefa
export const deleteTask = async (id: string): Promise<boolean> => {
  try {
    const response = await fetch(`${API_BASE_URL}/tasks/${id}`, {
      method: "DELETE",
    });
    if (!response.ok) {
      throw new Error(`Erro ao deletar tarefa: ${response.statusText}`);
    }
    return true;
  } catch (error) {
    console.error("Erro ao deletar tarefa:", error);
    return false;
  }
};

// Função auxiliar para filtrar tarefas por status
export const getTasksByStatus = (tasks: Task[], status: ColumnType): Task[] => {
  return tasks.filter((task) => task.status === status);
};

// Função auxiliar para obter tarefas pendentes
export const getPendingTasks = (tasks: Task[]): Task[] => {
  return getTasksByStatus(tasks, "todo");
};

// Função auxiliar para obter tarefas em progresso
export const getInProgressTasks = (tasks: Task[]): Task[] => {
  return getTasksByStatus(tasks, "in_progress");
};

// Função auxiliar para obter tarefas concluídas
export const getCompletedTasks = (tasks: Task[]): Task[] => {
  return getTasksByStatus(tasks, "done");
};

// Função para alternar status de uma tarefa (compatibilidade)
export const toggleTaskStatus = (tasks: Task[], taskId: string): Task[] => {
  return tasks.map((task) =>
    task.id === taskId ? { ...task, completed: !task.completed } : task
  );
};
