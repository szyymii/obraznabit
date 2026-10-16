def obraz_na_binarny(sciezka_obrazu, sciezka_wyjsciowa_bin):
	#wczytanie obrazu
	with open(sciezka_obrazu, 'rb') as plik_obrazu:
		dane_bajtowe = plik_obrazu.read()

	#zamiana bajtow na znaki 0 i 1 
	ciag_bitow = ''.join(format(bajt, '08b') for bajt in dane_bajtowe)

	#zapis txt lmao
	with open(sciezka_wyjsciowa_bin, 'w', encoding='utf-8') as plik_binarny:
		plik_binarny.write(ciag_bitow)

	print(f'Obraz zapisany jako kod binarny')
	return dane_bajtowe

def binarny_na_obraz(sciezka_binarna, sciezka_wyjsciowa_obrazu):
	#once upon a time wczytanie txt z 01
	with open(sciezka_binarna, 'r', encoding='utf-8') as plik_binarny:
		ciag_bitow = plik_binarny.read()

	#z ciagu na bajty once again...
	bajt_lista = bytearray()
	for i in range(0, len(ciag_bitow), 8):
		bajt_str = ciag_bitow[i : i + 8]
		if len(bajt_str) == 8:
			bajt_lista.append(int(bajt_str, 2))

	#zapis do obrazu znowu XDD
	with open(sciezka_wyjsciowa_obrazu, 'wb') as plik_obrazu:
		plik_obrazu.write(bajt_lista)

	print(f'Obraz zwrocony lmao: {sciezka_wyjsciowa_obrazu}')

if __name__ == '__main__':
	nazwa_oryginalu = 'obraz.jpg' #nazwa obrazu z tego samego folderu co ten plik
	plik_tymczasowy_bin = 'obraz_binarny.txt'
	nazwa_odtworzonego = 'odtworzone_zdjecie.jpg'

	#zamiana obrazu na ciag bitow 0 i 1
	obraz_na_binarny(nazwa_oryginalu, plik_tymczasowy_bin)

	#zamiana ciagu na obraz once again... jebany kurwa trans
	binarny_na_obraz(plik_tymczasowy_bin, nazwa_odtworzonego)