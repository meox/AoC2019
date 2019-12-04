defmodule Day03.Parser do
  def parse(file_name) do
    file_name
    |> read_file()
    |> String.split("\n")
    |> Enum.filter(fn line -> line != "" end)
    |> Enum.map(fn wire ->
      wire
      |> String.split(",")
      |> Enum.filter(fn token -> token != "" end)
      |> process()
    end)
  end

  defp read_file(file_name) do
    {:ok, content} = File.read(file_name)
    content
  end

  defp process(list) do
    list
    |> Enum.map(fn
      "L" <> r -> {:left, String.to_integer(r)}
      "R" <> r -> {:right, String.to_integer(r)}
      "U" <> r -> {:up, String.to_integer(r)}
      "D" <> r -> {:down, String.to_integer(r)}
    end)
  end
end
