import 'package:flutter/material.dart';
import 'package:http/http.dart' as http;
import 'dart:convert';
import 'barang.dart'; // Pastikan path ini benar

void main() {
  runApp(const MaterialApp(
    title: "Contoh",
    home: Home(),
  ));
}

class Home extends StatefulWidget {
  const Home({super.key});

  @override
  State<Home> createState() => _HomeState();
}

class _HomeState extends State<Home> {
  List<Barang> listProduct = [];
  bool loading = false;

  Future<void> __FetchDataProduct() async {
    String url = "http://10.0.2.2:911/barang";
    final response = await http.get(Uri.parse(url));

    if (response.statusCode == 200) {
      setState(() {
        loading = true;
        final responseData = jsonDecode(response.body);

        print('Response Data: $responseData');

        // Pastikan 'Data' ada dan bukan null
        if (responseData['Data'] != null) {
          final data = responseData['Data'] as List<dynamic>;

          listProduct = data.map((item) => Barang.fromJson(item as Map<String, dynamic>)).toList();
        }

        loading = false;
      });
    } else {
      // Tangani kasus ketika status code bukan 200
      print('Failed to load data');
    }
  }

  @override
  void initState() {
    super.initState();
    __FetchDataProduct();
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: const Text("Contoh"),
        backgroundColor: Colors.blue[200],
      ),
      body: Container(
        child: loading
            ? const Center(child: CircularProgressIndicator())
            : ListView.builder(
                itemCount: listProduct.length,
                itemBuilder: (context, i) {
                  final getdata = listProduct[i];
                  return Container(
                    child: ItemList(kodebarang: getdata.kodebarang, namabarang: getdata.namabarang) 
                  );
                },
              ),
      ),
    );
  }
}

class ItemList extends StatefulWidget {
  final String namabarang;
  final String kodebarang;

  const ItemList({super.key, required this.kodebarang, required this.namabarang});

  @override
  State<ItemList> createState() => _ItemListState();
}

class _ItemListState extends State<ItemList> {
  @override
  Widget build(BuildContext context) {
    return Container(
      child: Column(
        children: <Widget>[
          Card(
            child: ListTile(
              title: Text(widget.namabarang),
              subtitle: Text(widget.kodebarang)
            ),
          )
        ],
      ),
    );
  }
}
