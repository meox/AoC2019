defmodule Day03.WireController do
  use GenServer

  @me __MODULE__

  def start_link() do
    GenServer.start_link(@me, {%{}, []})
  end

  @impl true
  def init(state) do
    {:ok, state}
  end

  ### API

  @spec position(pid(), atom(), number(), number(), number()) :: any()
  def position(controller, direction, steps, x, y) do
    GenServer.cast(controller, {:pos, direction, self(), x, y, steps})
  end

  def matches(controller) do
    GenServer.call(controller, :get_matches)
  end

  ### CALLBACK

  @impl true
  def handle_call(:get_matches, _from, {_pos_map, matches} = state) do
    {:reply, matches, state}
  end

  @impl true
  def handle_cast({:pos, dir, id, x, y, steps}, {pos_map, matches}) do
    new_state =
      pos_map
      |> Map.update({x, y}, [{id, dir, steps}], fn l -> [{id, dir, steps} | l] end)
      |> update_state(matches, x, y)

    {:noreply, new_state}
  end

  defp update_state(pos_map, matches, x, y) do
    new_mathces =
      pos_map
      |> Map.get({x, y})
      |> update_matches(matches, {x, y})

    {pos_map, new_mathces}
  end

  defp update_matches([{id_a, dir_a, steps_a}, {id_b, dir_b, steps_b}], matches, {x, y}) when id_a != id_b and dir_a != dir_b do
    [{x, y, steps_a + steps_b} | matches]
  end
  defp update_matches(_, matches, _point), do: matches
end
