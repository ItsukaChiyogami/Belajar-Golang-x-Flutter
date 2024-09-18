class Barang {
  int id;
  String kodebarang;
  String namabarang;

  Barang({
    required this.id,
    required this.kodebarang,
    required this.namabarang,
  });

  factory Barang.fromJson(Map<String, dynamic> json) => Barang(
    id: json["Id"] != null ? json["Id"] as int : 0, // Default to 0 if null
    kodebarang: json["KodeBarang"],
    namabarang: json["NamaBarang"],
  );

  Map<String, dynamic> toJson() => {
    "Id": id,
    "KodeBarang": kodebarang,
    "NamaBarang": namabarang,
  };
}
