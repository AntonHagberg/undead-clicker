package main

import (
	"fmt"
	"time"
)

type player struct {
	biomass int
	souls int
}


type worker struct {
	name string
	lore string
	amount int
	production int
	baseCost int
	cost func(int, int) int


}

func dayWork(worker *worker, chBioMass chan int) {
	pay := worker.amount * worker.production
	chBioMass <- pay
}

func main()  {
	necromancer := player{
		biomass:	0,
		souls:		0,
	}

	skeletonHunter := worker{
		name:		"Skeleton Hunter",
		lore:  		"A hunter to be sent out to hunt",
		amount: 	1,
		production:	1,
		baseCost:	5,
		cost:		func(amount int, baseCost int) int {
			return baseCost + amount * baseCost
		},

	}
	skeletonFarmer := worker{
		name:		"Skeleton Farmer",
		lore:  		"A mindless skeleton farmer capable of working 24/7",
		amount: 	0,
		production:	5,
		baseCost:	25,
		cost:		func(amount int, baseCost int) int {
			return baseCost + amount * baseCost
		},
	}
	skeletonDruid := worker{
		name:		"Skeleton Druid",
		lore:  		"A powerful skeleton mage with nature magic",
		amount: 	0,
		production:	25,
		baseCost:	100,
		cost:		func(amount int, baseCost int) int {
			return baseCost + amount * baseCost
		},
	}

	chBioMass := make(chan int)

	for {
		go dayWork(&skeletonHunter, chBioMass)
		go dayWork(&skeletonFarmer, chBioMass)
		go dayWork(&skeletonDruid, chBioMass)
		if necromancer.biomass > skeletonDruid.cost(skeletonDruid.amount, skeletonDruid.baseCost) {
			necromancer.biomass = necromancer.biomass - skeletonDruid.cost(skeletonDruid.amount, skeletonDruid.baseCost)
			skeletonDruid.amount = skeletonDruid.amount + 1
		}
		if necromancer.biomass > skeletonFarmer.cost(skeletonFarmer.amount, skeletonFarmer.baseCost) {
			necromancer.biomass = necromancer.biomass - skeletonFarmer.cost(skeletonFarmer.amount, skeletonFarmer.baseCost)
			skeletonFarmer.amount = skeletonFarmer.amount + 1
		}
		if necromancer.biomass > skeletonHunter.cost(skeletonHunter.amount, skeletonHunter.baseCost) {
			necromancer.biomass = necromancer.biomass - skeletonHunter.cost(skeletonHunter.amount, skeletonHunter.baseCost)
			skeletonHunter.amount = skeletonHunter.amount + 1
		}
		empty := false
		for empty == false {
			select {
			case add := <- chBioMass:
				necromancer.biomass = necromancer.biomass + add

			case <-time.After(time.Second):
				empty = true
				fmt.Printf("amount of biomass %v\n", necromancer.biomass)
				fmt.Printf("amount of hunters %v\n", skeletonHunter.amount)
				fmt.Printf("amount of farmers %v\n", skeletonFarmer.amount)
				fmt.Printf("amount of druids %v\n", skeletonDruid.amount)
			}
		}
	}
}