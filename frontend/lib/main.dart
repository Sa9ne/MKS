import 'package:flutter/material.dart';

void main() {
  runApp(CybertrixApp());
}

class CybertrixApp extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      title: 'Cybertrix Lab',
      debugShowCheckedModeBanner: false,
      initialRoute: '/',
      routes: {
        '/': (context) => HomePage(),
        '/portfolio': (context) => PortfolioPage(),
      },
    );
  }
}

class HomePage extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(title: Text("Cybertrix Lab")),
      body: SingleChildScrollView(
        child: Column(
          children: [
            Row(
              mainAxisAlignment: MainAxisAlignment.center,
              children: [
                Placeholder(fallbackHeight: 150, fallbackWidth: 150),
                SizedBox(width: 20),
                Text(
                  "Cybertrix Lab\nОбщее описание лаборатории",
                  style: TextStyle(fontSize: 18),
                ),
              ],
            ),
            SizedBox(height: 20),

            // Направления
            Text("Направления деятельности", style: TextStyle(fontSize: 20)),
            DirectionBlock(
              title: "Направление 1",
              description: "Описание направления 3",
            ),

            DirectionBlock(
              title: "Направление 2",
              description: "Описание направления 3",
            ),

            DirectionBlock(
              title: "Направление 3",
              description: "Описание направления 3",
            ),

            SizedBox(height: 20),

            // Применение технологий
            Container(
              padding: EdgeInsets.all(20),
              color: Colors.green[100],
              child: Text("Применение наших технологий"),
            ),

            SizedBox(height: 20),

            // Карусель проектов
            Container(
              height: 200,
              color: Colors.yellow[100],
              child: PageView(
                children: [
                  ProjectCard(title: "Проект 1"),
                  ProjectCard(title: "Проект 2"),
                  ProjectCard(title: "Проект 3"),
                ],
              ),
            ),

            SizedBox(height: 20),

            // Контакты
            Container(
              color: Colors.purple[100],
              padding: EdgeInsets.all(16),
              child: Text("Контактная информация"),
            ),
          ],
        ),
      ),
    );
  }
}

class DirectionBlock extends StatelessWidget {
  final String title;
  final String? description;

  const DirectionBlock({required this.title, this.description});

  @override
  Widget build(BuildContext context) {
    return Card(
      margin: EdgeInsets.symmetric(vertical: 8, horizontal: 24),
      child: ListTile(
        title: Text(title),
        subtitle: description != null ? Text(description!) : null,
      ),
    );
  }
}

class ProjectCard extends StatelessWidget {
  final String title;

  const ProjectCard({required this.title});

  @override
  Widget build(BuildContext context) {
    return Center(child: Text(title, style: TextStyle(fontSize: 18)));
  }
}

class PortfolioPage extends StatelessWidget {
  final List<Map<String, String>> projects = [
    {'title': 'Проект 1', 'desc': 'Описание проекта 1'},
    {'title': 'Проект 2', 'desc': 'Описание проекта 2'},
    {'title': 'Проект 3', 'desc': 'Описание проекта 3'},
    {'title': 'Проект 4', 'desc': 'Описание проекта 4'},
  ];

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(title: Text("Портфолио")),
      body: Padding(
        padding: EdgeInsets.all(16),
        child: GridView.count(
          crossAxisCount: 2,
          mainAxisSpacing: 20,
          crossAxisSpacing: 20,
          children: projects.map((proj) {
            return Container(
              padding: EdgeInsets.all(12),
              color: Colors.grey[200],
              child: Column(
                children: [
                  Placeholder(fallbackHeight: 100),
                  SizedBox(height: 10),
                  Text(
                    proj['title']!,
                    style: TextStyle(fontWeight: FontWeight.bold),
                  ),
                  SizedBox(height: 5),
                  Text(proj['desc']!, textAlign: TextAlign.center),
                  SizedBox(height: 10),
                  ElevatedButton(onPressed: () {}, child: Text("Подробнее")),
                ],
              ),
            );
          }).toList(),
        ),
      ),
    );
  }
}
