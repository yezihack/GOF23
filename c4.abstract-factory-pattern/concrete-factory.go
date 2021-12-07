package abstractfactorypattern

// 具体工厂 01
type NikeFactory struct {
}

func (f *NikeFactory) MakeShoe() IShoeProduct {
	return &NikeShoe{
		Shoe: Shoe{
			size:     250,
			meterial: "布料",
		},
	}
}
func (f *NikeFactory) MakeTrousers() ITrousersProduct {
	return &NikeTrousers{
		Trousers: Trousers{
			color: "白色",
			size:  255,
		},
	}
}

// 具体工厂 02
type AdidasFactory struct {
}

func (f *AdidasFactory) MakeShoe() IShoeProduct {
	return &NikeShoe{
		Shoe: Shoe{
			size:     260,
			meterial: "真皮",
		},
	}
}
func (f *AdidasFactory) MakeTrousers() ITrousersProduct {
	return &NikeTrousers{
		Trousers: Trousers{
			color: "黑色",
			size:  275,
		},
	}
}
