package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var p, n, k int

	input, _ := os.Open("INPUT.TXT")
	defer input.Close()
	scanner := bufio.NewScanner(input)
	scanner.Scan()
	line := scanner.Text()
	parts := strings.Fields(line)

	p, _ = strconv.Atoi(parts[0])
	n, _ = strconv.Atoi(parts[1])
	k, _ = strconv.Atoi(parts[2])
	comands := make([]string, p)

	for i := 0; i < p; i++ {
		scanner.Scan()
		comands[i] = scanner.Text()
	}

	scanner.Scan()
	line = scanner.Text()
	num := strings.Fields(line)

	count := make(map[string]int)
	check := make(map[string]bool)

	selected := 0

	for i := 0; i < p && selected < n; i++ {
		univ := comands[i]
		id := num[i]
		key := univ + "|" + id
		if count[univ] < k && !check[key] {
			fmt.Printf("%s #%s\n", univ, id)
			//	writer.WriteString(univ + " #" + id + "\n")
			count[univ]++
			check[key] = true
			selected++
		}

	}
}

/* я решил задачу, но доебуются к выводу, я хз что не так с ним 
#include <fstream>

#include <vector>

#include <string>

#include <map>



using namespace std;



struct Team {

    string name;

    size_t number;

};



int main() {

    ifstream in("input.txt");

    ofstream out("output.txt");



    size_t p, n, k;

    in >> p >> n >> k;

    in.get();                       // читаем конец строки \n

    vector <Team> teams(p);

    for (size_t i = 0; i < p; ++i) {

        getline(in, teams[i].name); // читаем имена команд

    }

    for (size_t i = 0; i < p; ++i) {

        in >> teams[i].number;      // и номера команд

    }

    size_t invitations = 0;         // количество приглашенных

    map<string, size_t> university_teams; // создаем словарь

    for (auto& team : teams) {            // и пробегаем по данным команд

        if (invitations >= n)             // приглашенных больше лимита числа команд от университета?

            break;

        auto it = university_teams.insert(make_pair(team.name, 0)).first;

        // итератор указывает на вновь созданную в словаре запись: название команды -> 0

        ++it->second;       // увеличиваем счетчик команд на 1

        if (it->second > k) // если вышли из лимита команд на университет - продолжаем цикл с начала

            continue;

        out << team.name << " #" << team.number << endl; // выводим команду по правилам

        ++invitations;  // увеличиваем счетчик приглашенных

    }

    return 0;

}

*/
